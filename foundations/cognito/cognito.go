package cognito

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	cip "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/golang-jwt/jwt/v4"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/ramsfords/backend/configs"
)

type CognitoClient struct {
	Conf   configs.AwsConfig
	Client *cip.Client
	Key    jwk.Set
}
type ConfirmEmailData struct {
	ClientId         string `json:"client_id"`
	UserName         string `json:"user_name"`
	ConfirmationCode string `json:"confirmation_code"`
}

func GetCognitoClient(conf *configs.Config) (CognitoClient, error) {
	cc := CognitoClient{}
	confs := aws.Config{
		Region:           conf.GetAwsConfig().CognitoRegion,
		Credentials:      conf,
		RetryMaxAttempts: 10,
	}
	cipCLient := cip.NewFromConfig(confs)
	cc.Client = cipCLient
	cc.Key = cc.FetchKeySet(context.Background())
	return cc, nil
}

func NewClient(conf *configs.Config) (CognitoClient, error) {
	ct := CognitoClient{}
	awsConf := configs.GetConfig().GetAwsConfig()
	ct.Conf = awsConf
	ct.Key = ct.FetchKeySet(context.Background())
	confs := aws.Config{
		Region:           conf.GetAwsConfig().CognitoRegion,
		Credentials:      conf,
		RetryMaxAttempts: 10,
	}
	cipCLient := cip.NewFromConfig(confs)
	ct.Client = cipCLient
	return ct, nil
}

func (cc *CognitoClient) FetchKeySet(ctx context.Context) jwk.Set {
	keySet, err := jwk.Fetch(ctx, cc.Conf.JWKUrl)
	if err != nil {
		return nil
	}
	return keySet
}

func (cc *CognitoClient) Validate(ctx context.Context, tokenStr string) (jwt.Token, error) {
	// JWT Parse - it's actually doing parsing, validation and returns back a token.
	// Use .Parse or .ParseWithClaims methods from https://github.com/dgrijalva/jwt-go
	tokeKey, err := jwt.ParseWithClaims(tokenStr, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {

		// Verify if the token was signed with correct signing method
		// AWS Cognito is using RSA256 in my case
		_, ok := token.Method.(*jwt.SigningMethodRSA)

		if !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// Get "kid" value from token header
		// "kid" is shorthand for Key ID
		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, errors.New("kid header not found")
		}

		// "kid" must be present in the public keys set
		keys, ok := cc.Key.LookupKeyID(kid)
		if !ok {
			return nil, fmt.Errorf("key %v not found", kid)
		}
		// In our case, we are returning only one key = keys[0]
		// Return token key as []byte{string} type
		var tokenKey interface{}
		if err := keys.Raw(&tokenKey); err != nil {
			return nil, errors.New("failed to create token key")
		}

		return tokenKey, nil
	})
	if err != nil {
		log.Println("ParseWithClaims : error occured validating cognito token", err)
	}

	return *tokeKey, nil
}
