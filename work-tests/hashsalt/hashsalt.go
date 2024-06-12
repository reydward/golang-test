package hashsalt

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/scrypt"
)

const (
	saltSize = 16
	hashSize = 32
	scryptN  = 16384
	scryptR  = 8
	scryptP  = 1
)

func HashKeyWithSalt(key string) (string, error) {
	salt := make([]byte, saltSize)
	_, err := rand.Read(salt)

	if err != nil {
		return "", err
	}

	scryptedKey, err := scrypt.Key([]byte(key), salt, scryptN, scryptR, scryptP, hashSize)
	if err != nil {
		return "", err
	}

	hashedString := base64.StdEncoding.EncodeToString(append(scryptedKey, salt...))

	return hashedString, nil
}

func CheckHashKey(key string, hashedString string) (bool, error) {
	hashedBytes, err := base64.StdEncoding.DecodeString(hashedString)

	if err != nil {
		return false, err
	}

	hashedKey := hashedBytes[:len(hashedBytes)-saltSize]
	salt := hashedBytes[len(hashedBytes)-saltSize:]

	scryptedKey, err := scrypt.Key([]byte(key), salt, scryptN, scryptR, scryptP, hashSize)
	if err != nil {
		return false, err
	}

	return bytes.Equal(scryptedKey, hashedKey), nil
}
