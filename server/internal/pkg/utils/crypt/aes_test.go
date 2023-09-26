// Package test
// ---------------------------------
// @file      : aes_test.go
// @project   : faceto-ai
// @author    : zhangxiubo
// @time      : 2023/6/7 14:32
// @desc      : file description
// ---------------------------------
package crypt

import (
	"crypto/md5"
	"encoding/hex"
	"faceto-ai/internal/pkg/utils/helper"
	"fmt"
	"testing"
)

func TestEncryptByAes(t *testing.T) {
	uuid := helper.Generator()
	fmt.Println(uuid)

	hash := md5.Sum([]byte(uuid))
	key := hex.EncodeToString(hash[:])
	fmt.Println(key)

	enc, err := EncryptByAes(key, key)
	if err != nil {
		fmt.Println("EncryptByAes Err", err)
		return
	}
	fmt.Println("encrypt:", enc)

	dec, err := DecryptByAes(enc, key)
	if err != nil {
		fmt.Println("DecryptByAes Err", err)
		return
	}
	fmt.Println("decrypt:", string(dec))
}

func TestEncryptByAes2(t *testing.T) {
	key := "d977307f-d829-433e-a42a-e7a339c5a016"
	hash := md5.Sum([]byte(key))
	cryptVal, err := EncryptByAes(key, hex.EncodeToString(hash[:]))
	if err != nil {
		fmt.Println("EncryptByAes Err", err)
		return
	}
	fmt.Println(cryptVal)
}
