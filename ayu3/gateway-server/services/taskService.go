package services

import (
	"log"
	"strconv"

	"github.com/ValeHenriquez/example-rabbit-go/gateway-server/internal"
	"github.com/ValeHenriquez/example-rabbit-go/gateway-server/models"
	"github.com/gofiber/fiber/v2"
)



func CreateTask(c *fiber.Ctx) error {
	actionType := "CREATE_TASK"
	return internal.HandleRPCRequest(0, c, actionType)
}

func GetTasks(c *fiber.Ctx) error {
	actionType := "GET_TASKS"
	return internal.HandleRPCRequest(0, c, actionType)
}

func GetTask(c *fiber.Ctx) error {
	actionType := "GET_TASK"
	strID := c.Params("id")
	id, err := strconv.ParseUint(strID, 10, 32)
	if err != nil {
		log.Printf("Failed to parse body: %s", err)
		return c.Status(400).JSON(models.Response{
			Success: "false",
			Message: "Bad Request",
		})
	}
	return internal.HandleRPCRequest(id, c, actionType)
}

func UpdateTask(c *fiber.Ctx) error {
	return c.SendString("Update Task")
}

func DeleteTask(c *fiber.Ctx) error {
	return c.SendString("Delete Task")
}