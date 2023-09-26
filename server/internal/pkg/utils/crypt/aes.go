package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

//Encryption process:
//  1. Process the data, pad the data using PKCS7 (when the key length is not enough, fill in the missing bits).
//  2. Encrypt the data using the CBC encryption mode in the AES encryption method.
//  3. Encrypt the obtained encrypted data using base64 to get a string.
// Decryption process is the opposite.

// If the string is 16, 24, or 32 bits long, it corresponds to the AES-128, AES-192, and AES-256 encryption methods, respectively.
// The key cannot be leaked.
var PwdKey = []byte("WITHCONTEXTAI-FACETOAI")

// pkcs7Padding padding
func pkcs7Padding(data []byte, blockSize int) []byte {
	//Determine the length of the missing bits. At least 1, at most blockSize.
	padding := blockSize - len(data)%blockSize
	//Fill in the missing bits. Copy []byte{byte(padding)} padding times.
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// pkcs7UnPadding The reverse operation of padding
func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("encrypted string error!")
	}
	//Get the number of padding bits
	unPadding := int(data[length-1])
	return data[:(length - unPadding)], nil
}

// AesEncrypt encryption
func AesEncrypt(data []byte, key []byte) ([]byte, error) {
	//Create an encryption instance
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//Determine the size of the encryption block
	blockSize := block.BlockSize()
	//Padding
	encryptBytes := pkcs7Padding(data, blockSize)
	//Initialize the encrypted data receiving slice
	crypted := make([]byte, len(encryptBytes))
	//Use CBC encryption mode
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	//Perform encryption
	blockMode.CryptBlocks(crypted, encryptBytes)
	return crypted, nil
}

// AesDecrypt decryption
func AesDecrypt(data []byte, key []byte) ([]byte, error) {
	//Create an instance
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//Get the size of the block
	blockSize := block.BlockSize()
	//Use CBC
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	//Initialize the decryption data receiving slice
	crypted := make([]byte, len(data))
	//Perform decryption
	blockMode.CryptBlocks(crypted, data)
	//Remove padding
	crypted, err = pkcs7UnPadding(crypted)
	if err != nil {
		return nil, err
	}
	return crypted, nil
}

// EncryptByAes Aes encryption and then base64 encoding
func EncryptByAes(data string, secret string) (string, error) {
	res, err := AesEncrypt([]byte(data), []byte(secret))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(res), nil
}

// DecryptByAes Aes decryption
func DecryptByAes(data string, secret string) (string, error) {
	dataByte, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	decByte, err := AesDecrypt(dataByte, []byte(secret))
	if err != nil {
		return "", err
	}
	return string(decByte), nil
}
