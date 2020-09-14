package utils

import (
	"math/rand"
	"time"
)

const str = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// GenRandomString generate random string
func GenRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	data := make([]byte, length)
	l := len(str)
	for i := 0; i < length; i++ {
		idx := rand.Intn(l - 1)
		data[i] = str[idx]
	}

	return string(data)
}
