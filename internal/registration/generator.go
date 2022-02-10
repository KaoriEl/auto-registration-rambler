package registration

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

func GenerateString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

var NumRunes = []rune("0123456789")

func GenerateNum(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = NumRunes[rand.Intn(len(NumRunes))]
	}
	return string(b)
}
