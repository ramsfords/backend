package vault

import (
	"errors"
	"sync"
)

type VendorName string
type TokenType string

type Vault struct {
	Mutex     sync.Mutex
	TokenBank map[VendorName]map[TokenType]string
}

const (
	RapidShipLTL VendorName = "rapid_ship_ltl"
	Schneider    VendorName = "schneider"
	AuthType     TokenType  = "auth_type"
)

func New() *Vault {
	vault := &Vault{}
	vault.TokenBank = make(map[VendorName]map[TokenType]string)
	return vault
}

func avoidNil(vault *Vault, vendorName VendorName) {
	if vault.TokenBank == nil {
		vault.TokenBank = make(map[VendorName]map[TokenType]string)
	}
	if vault.TokenBank[vendorName] == nil {
		vault.TokenBank[vendorName] = make(map[TokenType]string)
	}
}
func (vault *Vault) AddAuthToken(vendorName VendorName, token string) {
	vault.Mutex.Lock()
	defer vault.Mutex.Unlock()
	avoidNil(vault, vendorName)
	vault.TokenBank[vendorName][AuthType] = token
}
func (vault *Vault) GetAuthToken(vendorName VendorName) (string, error) {
	vault.Mutex.Lock()
	defer vault.Mutex.Unlock()
	avoidNil(vault, vendorName)
	token, ok := vault.TokenBank[vendorName][AuthType]
	if !ok {
		return "", errors.New("could not get keys for vendor")
	}
	return token, nil
}
