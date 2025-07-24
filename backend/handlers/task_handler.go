package handlers

import (
	"TaskManager/database"
	"TaskManager/models"

	"github.com/gofiber/fiber/v2"
)

func CreateTask(ctx *fiber.Ctx) error {
	var task models.Task
	if err := ctx.BodyParser(&task); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "Data can not verified",
			})
	}	
	database.DB.Create(&task)
	return ctx.Status(201).JSON(task)
}

func GetAllTask(ctx *fiber.Ctx) error {
	var tasks []models.Task
	if err := database.DB.Find(&tasks).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": "Error",
		})
	}
	return ctx.JSON(tasks)
}

func GetTask(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var task models.Task

	if err := database.DB.First(&task, id).Error; err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"error": err,
		})
	}
	return ctx.JSON(task)
}

func UpdateTask(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var task models.Task

	if err := database.DB.First(&task, id).Error; err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"error": "Can not find task",
		})
	}

	if err := ctx.BodyParser(&task); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "Data not verify",
		})
	}

	database.DB.Save(&task)
	return ctx.JSON(task)
}

func DeleteTask(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var task models.Task
	if err := database.DB.First(&task, id).Error; err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"error": "Can not find task",
		})
	}
	database.DB.Delete(&task)
	return ctx.SendStatus(fiber.StatusNoContent)
}