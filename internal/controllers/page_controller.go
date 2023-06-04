package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func UserPage(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "This is User",
	})
}

func AdminPage(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "This is Admin",
	})
}
