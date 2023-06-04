package middlewares

import (
	"errors"
	auth "go-auth-jwt/internal/authentication/jwt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func UserAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "request does not contain an access token",
			})
		}

		token, err := jwt.ParseWithClaims(tokenString, &auth.JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(auth.JwtKey), nil
		})
	
		if err != nil {
			return err
		}
	
		claims, ok := token.Claims.(*auth.JWTClaim)
		if !ok {
			err = errors.New("couldn't parse claims")
			return err
		}
	
		if claims.Role != "user" {
			return c.Status(http.StatusForbidden).JSON(fiber.Map{
				"error": "Forbidden",
			})
		}
	
		expirationTime := time.Unix(claims.ExpiresAt.Unix(), 0)
		if expirationTime.Before(time.Now()) {
			err = errors.New("token expired")
			return err
		}

		return c.Next()
	}
}
