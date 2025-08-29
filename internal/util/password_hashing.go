package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	// Cost 12 is recommended for production (takes ~250ms)
	// Cost 10 = ~60ms, Cost 14 = ~1s
	const cost = 12
	
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func VerifyPassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}