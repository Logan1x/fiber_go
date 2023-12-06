// routes/user.go
package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/logan1x/fiber_go/controller"
)

func UserRoutes(app *fiber.App) {
	user := app.Group("/user")

	user.Get("/", controller.UsersIndex)
	user.Get("/:id", controller.UsersShow)
	user.Post("/", controller.UsersCreate)
	user.Put("/:id", controller.UsersUpdate)
	user.Delete("/:id", controller.UsersDelete)
}
