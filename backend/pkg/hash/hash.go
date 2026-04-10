package hash

import "golang.org/x/crypto/bcrypt"

const defaultCost = bcrypt.DefaultCost

// HashPassword generates a bcrypt hash of the given password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), defaultCost)
	return string(bytes), err
}

// CheckPassword compares a plain-text password with a bcrypt hash.
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
