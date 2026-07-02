package routesV1

import (
	"my-diet-server/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupFriendsRoutes(router fiber.Router) {
	app := router.Group("/friends")

	app.Get("/", middleware.DeserializeUser)
}
