package routesV1

import (
	handlersV1 "my-diet-server/internal/handler/v1"
	"my-diet-server/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRecipesRoutes(router fiber.Router) {
	app := router.Group("/recipes")

	app.Get("/", middleware.DeserializeUser, handlersV1.GetRecipesList)
	app.Get("/:recipeId", middleware.DeserializeUser, handlersV1.GetRecipe)
	app.Delete("/:recipeId", middleware.DeserializeUser, handlersV1.DeleteRecipe)
	app.Patch("/:recipeId", middleware.DeserializeUser, handlersV1.UpdateRecipe)
	app.Post("/", middleware.DeserializeUser, handlersV1.CreateRecipe)
}
