package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/logan1x/fiber_go/initializers"
	"github.com/logan1x/fiber_go/models"
)

func UsersIndex(c *fiber.Ctx) error {
	// get all users

	var users []models.User

	result := initializers.DB.Find(&users)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot find users",
			"error":   result.Error,
		})
	}

	// return all users
	return c.Status(fiber.StatusOK).JSON(users)

}

func UsersShow(c *fiber.Ctx) error {
	return c.SendString("Users Show")
}

func UsersCreate(c *fiber.Ctx) error {

	// get request body

	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	// create user

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot create user",
			"error":   result.Error,
		})
	}

	// return user

	return c.Status(fiber.StatusCreated).JSON(user)

}

func UsersUpdate(c *fiber.Ctx) error {
	return c.SendString("Users Update")
}

func UsersDelete(c *fiber.Ctx) error {
	return c.SendString("Users Delete")
}
