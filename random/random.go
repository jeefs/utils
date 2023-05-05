package utils

import (
	"crypto/rand"
	"math/big"
)

var defaultPool = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandString(n int, pool []rune) string {
	if len(pool) == 0 {
		pool = defaultPool
	}
	b := make([]rune, n)
	for i := range b {
		l := len(pool)
		randN, _ := rand.Int(rand.Reader, big.NewInt(int64(l)))
		b[i] = pool[randN.Int64()]
	}
	return string(b)
}
