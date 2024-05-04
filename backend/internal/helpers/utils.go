package helpers

import (
	"crypto/rand"
	"log"
	"math/big"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func GenerateRandomString(n int) (string, error) {
	// Don't return bytes just return a secure string
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	bytes := make([]byte, n)
	for i := range bytes {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		bytes[i] = letters[n.Int64()]
	}
	return string(bytes), nil
}

func HashPassword(password string) string {
	pw := []byte(password)
	result, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err.Error())
	}
	return string(result)
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
