package router

import (
	handlersV1 "my-diet-server/internal/handler/v1"
	routesV1 "my-diet-server/internal/routes/v1"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	app.Static("/public/images", "./storage/images")

	api := app.Group("/api", logger.New())

	// Check server is online
	api.Get("/status", handlersV1.CheckServerStatus)

	routesV1.SetupAuthRoutes(api)
	routesV1.SetupUserRoutes(api)
	routesV1.SetupRecipesRoutes(api)
	routesV1.SetupMenuRoutes(api)
	routesV1.SetupProductsRoutes(api)
	routesV1.SetupImagesRoutes(api)
}
