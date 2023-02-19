package model

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"

	"github.com/ramsfords/backend/foundations/logger"
)

type Crypto struct {
	secretkey string        `json:"-"`
	logger    logger.Logger `json:"-"`
	bytes     []byte        `json:"-"`
}

func New(secretKey string) *Crypto {
	return &Crypto{
		secretkey: secretKey,
		bytes:     []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05},
	}
}

func encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func (crypto Crypto) getBytes(data interface{}) []byte {
	// create json bytes
	bytes, err := json.Marshal(data)
	if err != nil {
		crypto.logger.Error(err)
		return nil
	}
	return bytes
}

// Encrypt method is to encrypt or hide any classified text
func (crypto Crypto) Encrypt(data interface{}) (string, error) {
	block, err := aes.NewCipher([]byte(crypto.secretkey))
	if err != nil {
		crypto.logger.Error(err)
		return "", err
	}
	plainBytes := crypto.getBytes(data)
	cfb := cipher.NewCFBEncrypter(block, crypto.bytes)
	plainBytes = make([]byte, len(plainBytes))
	cfb.XORKeyStream(plainBytes, plainBytes)
	return encode(plainBytes), nil
}

// Decrypt method is to extract back the encrypted text
func (crypto Crypto) Decrypt(data string) (string, error) {
	block, err := aes.NewCipher([]byte(crypto.secretkey))
	if err != nil {
		crypto.logger.Error(err)
		return "", err
	}
	cipherBytes := decode(data)
	cfb := cipher.NewCFBDecrypter(block, crypto.bytes)
	cipherBytes = make([]byte, len(cipherBytes))
	cfb.XORKeyStream(cipherBytes, cipherBytes)
	return string(cipherBytes), nil
}
