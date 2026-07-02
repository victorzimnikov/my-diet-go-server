package routesV1

import (
	handlersV1 "my-diet-server/internal/handler/v1"
	"my-diet-server/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupProductsRoutes(router fiber.Router) {
	app := router.Group("/products")

	app.Get("/", middleware.DeserializeUser, handlersV1.GetProductsList)
	app.Post("/", middleware.DeserializeUser, handlersV1.CreateProduct)
	app.Patch("/:productId", middleware.DeserializeUser, handlersV1.UpdateProduct)
	app.Get("/:productId", middleware.DeserializeUser, handlersV1.GetProduct)
	app.Delete("/:productId", middleware.DeserializeUser, handlersV1.DeleteProduct)
}
