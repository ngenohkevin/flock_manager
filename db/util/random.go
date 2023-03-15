package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomBreed() string {
	return RandomString(8)
}

func RandomProduction() int64 {
	return RandomInt(80, 20000)
}

func RandomHatchery() int64 {
	return RandomInt(80, 10000)
}

func RandomPremises() string {
	return RandomString(6)
}
