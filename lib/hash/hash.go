package hash

import (
	"golang.org/x/crypto/bcrypt"
)

func Generate(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func Compare(hashed, password string) error {
	err := bcrypt.CompareHashAndPassword(hashed, password)
	return err
}
