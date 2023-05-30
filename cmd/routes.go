package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kshyr/go-svelte-todo/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListTasks)
	app.Post("/task", handlers.CreateTask)
}
