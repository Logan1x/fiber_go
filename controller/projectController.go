package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/logan1x/fiber_go/initializers"
	"github.com/logan1x/fiber_go/models"
)

func ProjectsIndex(c *fiber.Ctx) error {
	// get all projects

	var projects []models.Project

	result := initializers.DB.Find(&projects)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot find projects",
			"error":   result.Error,
		})
	}

	// return all projects
	return c.Status(fiber.StatusOK).JSON(projects)
}

func ProjectsShow(c *fiber.Ctx) error {

	// get id from url params

	id := c.Params("id")

	project := new(models.Project)

	// check if project exists

	result := initializers.DB.First(&project, id)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot find project",
			"error":   result.Error,
		})
	}

	// return project

	return c.Status(fiber.StatusOK).JSON(project)
}

func ProjectsCreate(c *fiber.Ctx) error {

	// get request body

	project := new(models.Project)

	if err := c.BodyParser(project); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	// create project

	result := initializers.DB.Create(&project)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot create project",
			"error":   result.Error,
		})
	}

	// return project

	return c.Status(fiber.StatusCreated).JSON(project)

}

func ProjectsUpdate(c *fiber.Ctx) error {

	// get id from url params

	id := c.Params("id")

	project := new(models.Project)

	// check if project exists

	result := initializers.DB.First(&project, id)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot find project",
			"error":   result.Error,
		})
	}

	// get request body

	if err := c.BodyParser(project); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	// update project

	result = initializers.DB.Model(&project).Where("id = ?", id).Updates(models.Project{
		UserID:      project.UserID,
		ProjectName: project.ProjectName,
	})

	fmt.Print("Results : ", result)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot update project by given ID",
		})
	}

	// return project

	return c.Status(fiber.StatusOK).JSON(project)
}

func ProjectsDelete(c *fiber.Ctx) error {
	return c.SendString("Projects Delete")
}
