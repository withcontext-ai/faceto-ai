package crypt

import (
	"fmt"
	"testing"
)

func TestRSAGenKey(t *testing.T) {
	err := RSAGenKey(4096)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("gen success！")
}

func TestCryptRSA(t *testing.T) {
	str := "hello world!"
	fmt.Println("crypt before：", str)

	data, err := EncryptByRSA(str, "publicKey.pem")
	if err != nil {
		fmt.Println("EncryptByRSA Err", err)
		return
	}
	fmt.Println("crypt data：", data)

	data, err = DecryptByRSA(data, "privateKey.pem")
	if err != nil {
		fmt.Println("DecryptByRSA Err", err)
		return
	}
	fmt.Println("crypt after：", data)
}
