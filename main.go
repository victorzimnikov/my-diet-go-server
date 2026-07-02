package main

import (
	"fmt"
	"my-diet-server/config"
	"my-diet-server/database"
	"my-diet-server/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	// Middlewares
	app.Use(logger.New())
	app.Use(compress.New())
	app.Use(recover.New())
	// app.Use(limiter.New())

	// Connect to the Database
	database.ConnectDB()

	// Setup the router
	router.SetupRoutes(app)

	// Listen on PORT
	err := app.Listen(fmt.Sprintf(":%s", config.Config("SERVER_PORT")))

	if err != nil {
		fmt.Println(err)
	}
}
