package routesV1

import (
	handlersV1 "my-diet-server/internal/handler/v1"
	"my-diet-server/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupImagesRoutes(router fiber.Router) {
	app := router.Group("/images")

	app.Post("/upload", middleware.DeserializeUser, handlersV1.UploadImage)
}
