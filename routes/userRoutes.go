// routes/user.go
package routes

import "github.com/gofiber/fiber/v2"

func UserRoutes(app *fiber.App) {
	user := app.Group("/user")
	user.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("User route")
	})
	// Add more user routes here
}
