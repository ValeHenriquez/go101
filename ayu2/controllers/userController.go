package controllers

import (
	db "github.com/ValeHenriquez/go101/ayu2/config"
	"github.com/ValeHenriquez/go101/ayu2/models"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	db.DB.Select("id, first_name, last_name, email").Find(&users)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    users,
	})
}

func GetUser(c *fiber.Ctx) error {
	userId := c.Params("userId")
	var user models.User
	db.DB.Select("id, first_name, last_name, email").Where("id = ?", userId).First(&user)

	//asociar user con tareas
	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)

	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "User not found"})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    user,
	})

}
