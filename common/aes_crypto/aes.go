package aes_crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

const AesCryptoKey = ""

// PKCS7Padding fills plaintext as an integral multiple of the block length
func PKCS7Padding(p []byte, blockSize int) []byte {
	pad := blockSize - len(p)%blockSize
	padtext := bytes.Repeat([]byte{byte(pad)}, pad)
	return append(p, padtext...)
}

// PKCS7UnPadding removes padding data from the tail of plaintext
func PKCS7UnPadding(p []byte) ([]byte, error) {
	l := len(p)
	paddLen := int(p[l-1])
	if paddLen > l {
		return nil, errors.New("data is not illegal")
	}
	return p[:(l - paddLen)], nil
}

func AESCBCEncrypt(p, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	p = PKCS7Padding(p, block.BlockSize())
	ciphertext := make([]byte, len(p))
	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	blockMode.CryptBlocks(ciphertext, p)
	return ciphertext, nil
}

func AESCBCDecrypt(c, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	plaintext := make([]byte, len(c))
	blockMode := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	blockMode.CryptBlocks(plaintext, c)
	return PKCS7UnPadding(plaintext)
}

// Base64AESCBCEncrypt encrypts data with AES algorithm in CBC mode and encoded by base64
func Base64AESCBCEncrypt(p, key []byte) (string, error) {
	c, err := AESCBCEncrypt(p, key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(c), nil
}

// Base64AESCBCDecrypt decrypts cipher text encoded by base64 with AES algorithm in CBC mode
func Base64AESCBCDecrypt(c string, key []byte) ([]byte, error) {
	oriCipher, err := base64.StdEncoding.DecodeString(c)
	if err != nil {
		return nil, err
	}
	p, err := AESCBCDecrypt(oriCipher, key)
	if err != nil {
		return nil, err
	}
	return p, nil
}
