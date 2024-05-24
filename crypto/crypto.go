package crypto

import (
	"crypto/md5"
	"fmt"
	"github.com/mervick/aes-everywhere/go/aes256"
)

// 初次加密-MD5
func TentativeEncrypt(text string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(text)))
}

// 加严加密-
func TightenEncrypt(msg string) (string, error) {
	hexText := aes256.Encrypt(msg, "1234567812345678")
	return hexText, nil
}

// 解密
func Decrypt(cipherText string) (string, error) {
	plaintext := aes256.Decrypt(cipherText, "1234567812345678")
	return plaintext, nil
}
