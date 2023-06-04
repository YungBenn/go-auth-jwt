package bcrypt

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(passwors string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwors), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return string(hashedPassword), nil
}
