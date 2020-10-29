package crypto

import "testing"

func TestMD5(t *testing.T) {
	md5 := MD5("hello")
	t.Log(md5)

	sha256 := SHA256("world")
	t.Log(sha256)

	hashedText, _ := BcryptPwd("bjw")
	t.Log(hashedText, len(hashedText))

	flag := BcryptCompare(hashedText, "bjw")
	t.Log(flag)
}
