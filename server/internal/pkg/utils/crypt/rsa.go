package crypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"os"
)

func RSAGenKey(bits int) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}

	privateStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block1 := pem.Block{
		Type:  "private key",
		Bytes: privateStream,
	}
	fPrivate, err := os.Create("privateKey.pem")
	if err != nil {
		return err
	}
	defer fPrivate.Close()
	err = pem.Encode(fPrivate, &block1)
	if err != nil {
		return err
	}

	publicKey := privateKey.PublicKey
	publicStream, err := x509.MarshalPKIXPublicKey(&publicKey)
	//publicStream:=x509.MarshalPKCS1PublicKey(&publicKey)
	block2 := pem.Block{
		Type:  "public key",
		Bytes: publicStream,
	}
	fPublic, err := os.Create("publicKey.pem")
	if err != nil {
		return err
	}
	defer fPublic.Close()
	pem.Encode(fPublic, &block2)
	return nil
}

func EncryptByRSA(src string, path string) (res string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	fileInfo, _ := f.Stat()
	b := make([]byte, fileInfo.Size())
	f.Read(b)
	block, _ := pem.Decode(b)

	keyInit, err := x509.ParsePKIXPublicKey(block.Bytes)
	//keyInit1,err:=x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return
	}

	pubKey := keyInit.(*rsa.PublicKey)
	data, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, []byte(src))
	res = base64.StdEncoding.EncodeToString(data)
	return
}

func DecryptByRSA(src string, path string) (res string, err error) {
	decodeData, err := base64.StdEncoding.DecodeString(src)
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	fileInfo, _ := f.Stat()
	b := make([]byte, fileInfo.Size())
	f.Read(b)
	block, _ := pem.Decode(b)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	data, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, decodeData)
	res = string(data)
	return
}
