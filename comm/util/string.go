package util

import (
	"math/rand"
)

func RandomString(length int) string {
	const charset = "123456789abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
