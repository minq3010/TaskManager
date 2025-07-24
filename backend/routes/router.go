package routes

import (
	"TaskManager/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("tasks")

	api.Get("/", handlers.GetAllTask)
	api.Get("/:id", handlers.GetTask)
	api.Post("/", handlers.CreateTask)
	api.Put("/:id", handlers.UpdateTask)
	api.Delete("/:id", handlers.DeleteTask)
}