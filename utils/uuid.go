package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
)

func Md5(str string) string {
	tmp := md5.New()
	tmp.Write([]byte(str))
	return hex.EncodeToString(tmp.Sum(nil))
}

func Uuid() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return Md5(base64.URLEncoding.EncodeToString(b))
}
