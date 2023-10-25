package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const characters = "abcdefghijklmnopqrstuvwxyz1234567890-=!@#$%^&*()P{}:<>?ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const cif = "0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomNameOrEmail generates a random name or a random email if the isEmail bool is true of length n
func RandomNameOrEmail(n int, isEmail bool) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	if isEmail {
		sb.WriteString("@ex.com")
	}

	return sb.String()
}

// RandomPass generates a random password of length n
func RandomPass(n int) string {
	var sb strings.Builder
	k := len(characters)

	for i := 0; i < n; i++ {
		c := characters[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomPhoneNumber generates a random phone number of length n
func RandomPhoneNumber(n int) string {
	var sb strings.Builder
	k := len(cif)

	for i := 0; i < n; i++ {
		c := cif[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}
