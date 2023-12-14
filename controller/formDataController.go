package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/logan1x/fiber_go/initializers"
	"github.com/logan1x/fiber_go/models"
)

func FormDataIndex(c *fiber.Ctx) error {

	// get all formData

	var formData []models.FormData

	result := initializers.DB.Find(&formData)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot find formData",
			"error":   result.Error,
		})
	}

	// return all formData
	return c.Status(fiber.StatusOK).JSON(formData)
}

func FormDataShow(c *fiber.Ctx) error {

	// get id from url params

	id := c.Params("id")

	formData := new(models.FormData)

	// check if formData exists

	result := initializers.DB.First(&formData, id)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot find formData",
			"error":   result.Error,
		})
	}

	// return formData

	return c.Status(fiber.StatusOK).JSON(formData)
}

func FormDataCreate(c *fiber.Ctx) error {

	// get request body

	formData := new(models.FormData)

	if err := c.BodyParser(formData); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
			"error":   err,
		})

	}

	// create formData

	result := initializers.DB.Create(&formData)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot create formData",
			"error":   result.Error,
		})
	}

	// return formData

	return c.Status(fiber.StatusCreated).JSON(formData)
}

func FormDataCreateWithProject(c *fiber.Ctx) error {

	// get request body

	projectId := c.Params("projectId")

	formData := new(models.FormData)

	if err := c.BodyParser(formData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	// create formData

	formData.ProjectID = uuid.MustParse(projectId)

	result := initializers.DB.Create(&formData)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot create formData",
			"error":   result.Error,
		})
	}

	// return formData

	return c.Status(fiber.StatusCreated).JSON(formData)

}
func FormDataUpdate(c *fiber.Ctx) error {

	// get id from url params

	id := c.Params("id")

	formData := new(models.FormData)

	// check if formData exists

	result := initializers.DB.First(&formData, id)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot find formData",
			"error":   result.Error,
		})
	}

	// get request body

	if err := c.BodyParser(formData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	// update formData

	result = initializers.DB.Save(&formData)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot update formData",
			"error":   result.Error,
		})
	}

	// return formData

	return c.Status(fiber.StatusOK).JSON(formData)
}

func FormDataDelete(c *fiber.Ctx) error {

	// get id from url params

	id := c.Params("id")

	formData := new(models.FormData)

	// check if formData exists

	result := initializers.DB.First(&formData, id)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot find formData",
			"error":   result.Error,
		})
	}

	// delete formData

	result = initializers.DB.Delete(&formData)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot delete formData",
			"error":   result.Error,
		})
	}

	// return success message

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
	})
}
