package crypto

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

func MD5(text string) string {
	hash := md5.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

func SHA256(text string) string {
	hash := sha256.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

func BcryptPwd(text string) (string, error) {
	// 长度为 60 且每次都不一样的加密字符串
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashPwd), nil
}

func BcryptCompare(hashedPwd string, originPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(originPwd))
	if err != nil {
		return false
	}
	// 没有错误则密码匹配
	return true
}