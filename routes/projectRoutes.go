// routes/project.go
package routes

import "github.com/gofiber/fiber/v2"

func ProjectRoutes(app *fiber.App) {
	project := app.Group("/project")
	project.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Project route")
	})
	// Add more project routes here
}
