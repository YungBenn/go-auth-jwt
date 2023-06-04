package main

import (
	"go-auth-jwt/config"
	"go-auth-jwt/internal/database"
	"go-auth-jwt/internal/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func buildServer() error {
	env, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	db := database.Connect(env.DBurl)

	app := fiber.New()

	app.Use(logger.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK!")
	})

	routes.SetupRoutes(app, db)

	port := env.Port
	app.Listen(":" + port)

	return nil
}

func main() {
	err := buildServer()
	if err != nil {
		log.Fatal(err)
	}
}