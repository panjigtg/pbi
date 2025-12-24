package utils

import (
	"golang.org/x/crypto/bcrypt"
)

const passwordCost = bcrypt.DefaultCost

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		passwordCost,
	)

	return string(hash), err
}

func VerifyPassword(password, hashed string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashed),
		[]byte(password),
	)
}