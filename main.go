package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/logan1x/fiber_go/initializers"
	"github.com/logan1x/fiber_go/migrate"
	"github.com/logan1x/fiber_go/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	migrate.Migrate()
}

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to Fiber!")
}

func main() {
	app := fiber.New()

	app.Get("/", welcome)

	// Register routes
	routes.UserRoutes(app)
	routes.ProjectRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //default port
	}

	app.Listen(":" + port)
}
