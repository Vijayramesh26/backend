package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword - only for creating a new password / storing in DB
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CheckPassword - compares plaintext with stored hash
func CheckPassword(storedHash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
	return err == nil
}
