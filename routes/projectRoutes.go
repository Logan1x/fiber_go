// routes/project.go
package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/logan1x/fiber_go/controller"
)

func ProjectRoutes(app *fiber.App) {
	project := app.Group("/project")

	project.Get("/", controller.ProjectsIndex)
	project.Get("/:id", controller.ProjectsShow)
	project.Post("/", controller.ProjectsCreate)
	project.Put("/:id", controller.ProjectsUpdate)
	project.Delete("/:id", controller.ProjectsDelete)

	// Add more project routes here
}
