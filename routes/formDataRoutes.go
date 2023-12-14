package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/logan1x/fiber_go/controller"
)

func FormDataRoutes(app *fiber.App) {
	formData := app.Group("/formData")

	formData.Get("/", controller.FormDataIndex)
	formData.Get("/:id", controller.FormDataShow)
	formData.Post("/", controller.FormDataCreate)
	formData.Put("/:id", controller.FormDataUpdate)
	formData.Delete("/:id", controller.FormDataDelete)

	// Add more formData routes here
}
