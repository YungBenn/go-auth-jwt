package controllers

import (
	"go-auth-jwt/internal/authentication/bcrypt"
	"go-auth-jwt/internal/authentication/jwt"
	"go-auth-jwt/internal/entity"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RegisterDTO struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserController struct {
	DB *gorm.DB
}

func (u *UserController) Register(c *fiber.Ctx) error {
	var input RegisterDTO

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	password, err := bcrypt.HashPassword(input.Password)
	if err != nil {
		return err
	}

	user := entity.User{
		Name:     input.Name,
		Username: input.Username,
		Email:    input.Email,
		Password: password,
		Role:     input.Role,
	}

	result := u.DB.Save(&user)
	if result.Error != nil {
		return result.Error
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status":  201,
		"message": "New user created successfully",
	})
}

func (u *UserController) Login(c *fiber.Ctx) error {
	var input LoginDTO
	var user entity.User

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	record := u.DB.Where("username = ?", input.Username).First(&user)
	if record.Error != nil {
		return record.Error
	}

	checkPass := bcrypt.CheckPassword(input.Password, user.Password)
	if checkPass != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Wrong Password!",
		})
	}

	tokenString, err := jwt.GenerateJWT(input.Username, user.Role)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":   200,
		"message":  "Login Success",
		"username": user.Username,
		"token":    tokenString,
	})
}
