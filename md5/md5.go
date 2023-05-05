package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// 32
func MD5_32(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

// 16
func MD5_16(v string) string {
	return MD5_32(v)[8:24]
}
