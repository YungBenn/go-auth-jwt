package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte("supersecretkey")

type JWTClaim struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(username string, role string) (tokenString string, err error) {
	expirationTime := jwt.NewNumericDate(time.Now().Add(1 * time.Hour))
	claims := &JWTClaim{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expirationTime,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(JwtKey)

	return
}
