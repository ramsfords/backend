package keystore

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"path"
	"strings"
	"sync"

	"github.com/golang-jwt/jwt/v4"
	errs "github.com/ramsfords/backend/foundations/error"
)

var KeyStor *KeyStore

// KeyStore represents an in memory store implementation of the
// keystorer interface for use with the auth package
type KeyStore struct {
	mu    sync.RWMutex
	store map[string]*rsa.PrivateKey
}

// New construct an empty KeyStore ready for use

func New() *KeyStore {
	if KeyStor == nil {
		*KeyStor = KeyStore{
			store: make(map[string]*rsa.PrivateKey)}
		return KeyStor
	}
	return KeyStor

}

// NewMap construct a KeyStore with an initail set of keys.

func NewMap(store map[string]*rsa.PrivateKey) *KeyStore {
	KeyStor.store = store
	return KeyStor
}

// NewFS constructs a KeyStore based on a set of PEM files rooted inside
// of a directory. The name of each PEM file will be used as the key id.
// Example: keystore.NewFS(os.DirFS("/zarf/keys/"))
// Example: /config/keys/54bb2165-71e1-41a6-af3e-7da4a0e1e2c1.pem
func NewFS(fsys fs.FS) *KeyStore {
	ks := KeyStore{
		store: make(map[string]*rsa.PrivateKey),
	}
	fn := func(fileName string, dirEntry fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(errs.NewApiError(errs.ErrStartingComponents.Cod, fmt.Errorf("walkdir failure: %w", err).Error(), fmt.Errorf("walkdir failure: %w", err)))
		}
		if dirEntry.IsDir() {
			return nil
		}
		if path.Ext(fileName) != ".pem" {
			return nil
		}
		file, err := fsys.Open(fileName)
		if err != nil {
			log.Fatal(errs.NewApiError(errs.ErrStartingComponents.Code(), fmt.Errorf("opening key file: %w", err).Error(), fmt.Errorf("opening key file: %w", err)))
		}
		defer file.Close()
		// limit PEM file size to 1 megabyte. This should be reasonable for
		// almost any PEM file and prevents shenanigans like linking the file
		// to /dev/random or something like that.
		privatePEM, err := io.ReadAll(io.LimitReader(file, 1024*1024))
		if err != nil {
			log.Fatal(errs.NewApiError(errs.ErrStartingComponents.Code(), fmt.Errorf("reading auth private key: %w", err).Error(), fmt.Errorf("reading auth private key: %w", err)))
		}
		privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privatePEM)
		if err != nil {
			log.Fatal(errs.NewApiError(errs.ErrStartingComponents.Code(), fmt.Errorf("parsing auth private key: %w", err).Error(), fmt.Errorf("parsing auth private key: %w", err)))
		}

		ks.store[strings.TrimSuffix(dirEntry.Name(), ".pem")] = privateKey
		return nil

	}
	if err := fs.WalkDir(fsys, ".", fn); err != nil {
		log.Fatal(errs.NewApiError(errs.ErrStartingComponents.Code(), fmt.Errorf("walking directory: %w", err).Error(), fmt.Errorf("walking directory: %w", err)))
	}
	KeyStor = &ks
	return KeyStor

}

// Add adds a private key and combination kid to the store.
func (ks *KeyStore) Add(privateKey *rsa.PrivateKey, kid string) {
	ks.mu.Lock()
	defer ks.mu.Unlock()

	ks.store[kid] = privateKey
}

// Remove removes a private key and combination kid to the store.
func (ks *KeyStore) Remove(kid string) {
	ks.mu.Lock()
	defer ks.mu.Unlock()

	delete(ks.store, kid)
}

// PrivateKey searches the key store for a given kid and returns
// the private key.
func (ks *KeyStore) PrivateKey(kid string) (*rsa.PrivateKey, error) {
	ks.mu.RLock()
	defer ks.mu.RUnlock()

	privateKey, found := ks.store[kid]
	if !found {
		return nil, errors.New("kid lookup failed")
	}
	return privateKey, nil
}

// PublicKey searches the key store for a given kid and returns
// the public key.
func (ks *KeyStore) PublicKey(kid string) (*rsa.PublicKey, error) {
	ks.mu.RLock()
	defer ks.mu.RUnlock()

	privateKey, found := ks.store[kid]
	if !found {
		return nil, errors.New("kid lookup failed")
	}
	return &privateKey.PublicKey, nil
}

// Gets the random keys for the jwt creation
func (ks *KeyStore) Getkey() string {
	for key, _ := range ks.store {
		return key
	}
	return ""
}
