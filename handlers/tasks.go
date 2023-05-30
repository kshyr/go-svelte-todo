package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kshyr/go-svelte-todo/db"
	"github.com/kshyr/go-svelte-todo/models"
)

func ListTasks(c *fiber.Ctx) error {
	tasks := []models.Task{}

	db.DB.DB.Find(&tasks)

	return c.Status(200).JSON(tasks)
}

func CreateTask(c *fiber.Ctx) error {
	task := new(models.Task)

	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	db.DB.DB.Create(&task)

	return c.Status(200).JSON(task)
}
