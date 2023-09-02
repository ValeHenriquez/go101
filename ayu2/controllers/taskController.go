package controllers

import (
	db "github.com/ValeHenriquez/go101/ayu2/config"
	"github.com/ValeHenriquez/go101/ayu2/models"
	"github.com/gofiber/fiber/v2"
)

func GetTasks(c *fiber.Ctx) error {
	var tasks []models.Task
	db.DB.Select("id, title, description, done, user_id").Find(&tasks)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    tasks,
	})
}

func GetTask(c *fiber.Ctx) error {
	taskId := c.Params("taskId")
	var task models.Task
	db.DB.Select("id, title, description, done").Where("id = ?", taskId).First(&task)

	if task.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "Task not found"})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    task,
	})

}

func CreateTask(c *fiber.Ctx) error {
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Bad request", "error": err})
	}

	db.DB.Create(&task)

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    task,
	})
}

func UpdateTask(c *fiber.Ctx) error {
	taskId := c.Params("taskId")
	var task models.Task
	db.DB.Select("id, title, description, completed").Where("id = ?", taskId).First(&task)

	if task.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "Task not found"})
	}

	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Error", "error": err})
	}

	db.DB.Save(&task)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    task,
	})
}

func DeleteTask(c *fiber.Ctx) error {
	taskId := c.Params("taskId")
	var task models.Task
	db.DB.Select("id, title, description, completed").Where("id = ?", taskId).First(&task)

	if task.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "Task not found"})
	}

	db.DB.Delete(&task) //Change delete field to true and assing deleted_at
	//db.DB.Unscoped().Delete(&task) //Unscoped() is used to delete permanently

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
	})
}
