package routes

import (
	// "user-service/pkg/context"
	// "user-service/pkg/dto"
	"user-service/pkg/handlers"
	"user-service/pkg/jwt"
	"user-service/pkg/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App, userHandler *handlers.UserHandler, jwtSvc *jwt.JwtService) {

	api := app.Group("/api")
	api.Get("/swagger/*", swagger.New(swagger.Config{
		URL: "/swagger/doc.json", // The url pointing to API definition
	}))
	v1 := api.Group("/v1")
	v1.Post("/register",
		userHandler.PatientRegister)
	v1.Post("/login", userHandler.PatientLogin)
	v1.Use(middleware.JwtMiddleware(jwtSvc))
	v1.Get("/profile", userHandler.Profile)

}
