package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomString(length int) string {
	characters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	result := make([]rune, length)

	rand.Seed(time.Now().UnixNano()) // Rastgele sayı üreteci için tohumlama

	for i := range result {
		result[i] = characters[rand.Intn(len(characters))]
	}
	return string(result)
}
