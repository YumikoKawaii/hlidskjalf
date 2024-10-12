package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	cRand "crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
)

func HashString(input, hashKey string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum([]byte(hashKey)))
}

func Encrypt(input []byte, encryptKey string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(encryptKey))
	if err != nil {
		return nil, err
	}
	bytes := base64.StdEncoding.EncodeToString(input)
	cipherText := make([]byte, aes.BlockSize+len(bytes))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(cRand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cipherText[aes.BlockSize:], []byte(bytes))
	return cipherText, nil
}

func Decrypt(input []byte, encryptKey string) ([]byte, error) {
	inputBytes := []byte(input)
	block, err := aes.NewCipher([]byte(encryptKey))
	if err != nil {
		return nil, err
	}
	if len(inputBytes) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := inputBytes[:aes.BlockSize]
	inputBytes = inputBytes[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(inputBytes, inputBytes)
	data, err := base64.StdEncoding.DecodeString(string(inputBytes))
	if err != nil {
		return nil, err
	}
	return data, nil
}
