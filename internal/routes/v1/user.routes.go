package routesV1

import (
	handlersV1 "my-diet-server/internal/handler/v1"
	"my-diet-server/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router) {
	app := router.Group("/user")

	app.Get("/me", middleware.DeserializeUser, handlersV1.GetMe)

	// User friends routes
	SetupFriendsRoutes(app)
}
