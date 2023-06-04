package routes

import (
	"go-auth-jwt/internal/controllers"
	"go-auth-jwt/internal/middlewares"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	userController := controllers.UserController{
		DB: db,
	}

	api := app.Group("/api")

	// user
	api.Post("/register", userController.Register)
	api.Post("/login", userController.Login)

	// role user required
	userAuth := api.Group("/user").Use(middlewares.UserAuth())
	userAuth.Get("/", controllers.UserPage)

	// role admin required
	adminAuth := api.Group("/admin").Use(middlewares.AdminAuth())
	adminAuth.Get("/", controllers.AdminPage)
}
