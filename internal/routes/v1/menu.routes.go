package routesV1

import (
	"my-diet-server/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupMenuRoutes(router fiber.Router) {
	app := router.Group("menu")

	app.Get("/", middleware.DeserializeUser)
}
