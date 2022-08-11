package util

import (
	"math/rand"
	"strings"
	"time"
)

const alpabet = "abcdefghijklmnopqrstuvwxyz"
const number = "0123456789"
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomString(n int, x string) string {
	var sb strings.Builder
	k := len(x)

	for i := 0; i < n; i++ {
		c := x[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomName generates a random name
func RandomName() string {
	return RandomString(6, alpabet)
}

// RandomPassword generates a random password
func RandomPassword() string {
	return RandomString(4, charset)
}

// RandomPhone generates a random phone number
func RandomPhone() string {
	return RandomString(12, number)
}
