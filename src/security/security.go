package security

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(passwordString, passwordWithHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordWithHash), []byte(passwordString))
}
