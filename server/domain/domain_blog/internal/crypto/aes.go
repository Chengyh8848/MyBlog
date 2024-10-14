package crypto

import (
	"domain_blog/common"
	"encoding/base64"
	"fmt"
	"github.com/forgoer/openssl"
)

const (
	AesKey    = "3ef7fd1d5c124491"
	AesKeyApi = "web2ap#7d30bc5a1"
)

// Encrypt AES加密
func Encrypt(origData, key []byte) ([]byte, error) {
	dst, err := openssl.AesECBEncrypt(origData, key, openssl.PKCS7_PADDING)
	return dst, err
}

// Decrypt AES解密
func Decrypt(encrypted, key []byte) (string, error) {
	if len(encrypted) == 0 {
		return "", nil
	}
	dst, err := openssl.AesECBDecrypt(encrypted, key, openssl.PKCS7_PADDING)
	if err != nil {
		return "", err
	}
	return string(dst), nil
}

func EncryptData(data string) string {
	if data != "" {
		dataBytes, _ := Encrypt([]byte(data), []byte(AesKey))
		return base64.StdEncoding.EncodeToString(dataBytes)
	}
	return ""
}

func DecryptData(data string) (newData string) {
	defer func() {
		if err := recover(); err != nil {
			if common.Log != nil {
				common.Log.ErrorMsg("%v", err)
			}
			newData = ""
		}
	}()
	if data != "" {
		bytesPass, err := base64.StdEncoding.DecodeString(data)
		if err != nil {
			return ""
		}
		newData, err = Decrypt(bytesPass, []byte(AesKey))
		if err != nil {
			return ""
		}
		return
	}
	return ""
}

func EncryptApiData(data string) (string, error) {
	if data != "" {
		dataBytes, err := Encrypt([]byte(data), []byte(AesKeyApi))
		if err != nil {
			return "", err
		}
		return base64.URLEncoding.EncodeToString(dataBytes), nil
	}
	return "", fmt.Errorf("data err")
}

func DecryptApiData(data string) (string, error) {
	if data != "" {
		bytesPass, err := base64.URLEncoding.DecodeString(data)
		if err != nil {
			return "", err
		}
		newData, err := Decrypt(bytesPass, []byte(AesKeyApi))
		return newData, err
	}
	return "", fmt.Errorf("data err")
}
