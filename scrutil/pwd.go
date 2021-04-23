package scrutil

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 9)
	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}

func ValidateHashPassword(password, hashPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password)); err == nil {
		return true
	}
	return false
}
