package routesV1

import (
	handlersV1 "my-diet-server/internal/handler/v1"
	"my-diet-server/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(router fiber.Router) {
	app := router.Group("/auth")

	app.Post("/register", handlersV1.RegisterUser)
	app.Post("/login", handlersV1.Login)
	app.Post("/refresh", handlersV1.RefreshAccessToken)
	app.Get("/logout", middleware.DeserializeUser, handlersV1.LogoutUser)
}
