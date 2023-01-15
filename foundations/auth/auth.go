package auth

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v4"
)

var (
	ErrForbidden = errors.New("attempted action is not allowed")
)

// KeyLookup declares a method set of behavior for looking up
// private and public keys for JWT use.
type KeyLookup interface {
	PrivateKey(kid string) (*rsa.PrivateKey, error)
	PublicKey(kid string) (*rsa.PublicKey, error)
}

// Auth is used to authenticate clients. It can generate a token for a
// set of user claims and recreate the claims by parsing the token.
type Auth struct {
	activeKID string
	keyLookup KeyLookup
	method    jwt.SigningMethod
	parser    jwt.Parser
}

// KeyFunc is a function that will return one of the key(kid) from keystore
type keyFunc func() string

var KeyLookupFunction KeyLookup

func KeyProvider(t *jwt.Token) (interface{}, error) {
	kid, ok := t.Header["kid"]
	if !ok {
		log.Fatal(errors.New("missing key id (kid) in token header"))
	}
	kidID, ok := kid.(string)
	if !ok {
		return nil, errors.New("user token key id (kid) must be string")
	}
	return KeyLookupFunction.PublicKey(kidID)
}

// New creates an Auth to support authentication/authorization.
func New(kidProvider keyFunc, keyLookup KeyLookup) *Auth {
	KeyLookupFunction = keyLookup
	// The activeKID represents the private key used to signed new tokens.
	activeKID := kidProvider()
	_, err := keyLookup.PrivateKey(activeKID)
	if err != nil {
		log.Fatal(errors.New("active KID does not exist in store"))
	}

	method := jwt.GetSigningMethod("RS256")
	if method == nil {
		log.Fatal(errors.New("configuring algorithm RS256"))
	}

	// Create the token parser to use. The algorithm used to sign the JWT must be
	// validated to avoid a critical vulnerability:
	// https://auth0.com/blog/critical-vulnerabilities-in-json-web-token-libraries/
	parser := jwt.Parser{
		ValidMethods: []string{"RS256"},
	}

	return &Auth{
		activeKID: activeKID,
		keyLookup: keyLookup,
		method:    method,
		parser:    parser,
	}
}

// GenerateToken generates a signed JWT token string representing the user Claims.
func (a *Auth) GenerateToken(claims Claims) (string, error) {
	token := jwt.NewWithClaims(a.method, claims)
	token.Header["kid"] = a.activeKID

	privateKey, err := a.keyLookup.PrivateKey(a.activeKID)
	if err != nil {
		return "", errors.New("kid lookup failed")
	}

	str, err := token.SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("signing token: %w", err)
	}

	return str, nil
}

// ValidateToken recreates the Claims that were used to generate a token. It
// verifies that the token was signed using our key.
func (a *Auth) ValidateToken(tokenStr string) (Claims, bool) {
	var claims Claims
	token, err := a.parser.ParseWithClaims(tokenStr, &claims, KeyProvider)
	if err != nil {
		return Claims{}, false
	}

	if !token.Valid {
		return Claims{}, false
	}

	return claims, true
}
