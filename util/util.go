package util

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(text string) string  {
	password := []byte(text)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

func ComparePassword(hashedPassword , password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}
