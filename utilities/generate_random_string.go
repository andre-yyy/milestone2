package utility

import (
	"time"

	"golang.org/x/exp/rand"
)

func GenerateRandomString(length int) string {
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(uint64(time.Now().UnixNano()))
	randomString := make([]byte, length)
	for i := range randomString {
		randomString[i] = characters[rand.Intn(len(characters))]
	}
	return string(randomString)
}
