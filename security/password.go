package security

import "golang.org/x/crypto/bcrypt"

const DefaultCost = 10

// HashPassword converts plaintext password to hash
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), DefaultCost)
	return string(bytes), err
}

// VerifyPassword checks if password matches hash
func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
