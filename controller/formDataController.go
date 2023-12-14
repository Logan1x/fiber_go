package controller

import "github.com/gofiber/fiber/v2"

func FormDataIndex(c *fiber.Ctx) error {
	return c.SendString("FormData Index")
}

func FormDataShow(c *fiber.Ctx) error {
	return c.SendString("FormData Show")
}

func FormDataCreate(c *fiber.Ctx) error {
	return c.SendString("FormData Create")
}

func FormDataUpdate(c *fiber.Ctx) error {
	return c.SendString("FormData Update")
}

func FormDataDelete(c *fiber.Ctx) error {
	return c.SendString("FormData Delete")
}
