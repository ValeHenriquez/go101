package routes

import (
	"github.com/ValeHenriquez/example-rabbit-go/gateway-server/services"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("This is the Gateway API")
	})

	app.Post("/login", services.Login)
	app.Post("/register", services.SignUp)


	app.Post("createTask", services.CreateTask)
	app.Get("/tasks", services.GetTasks)
	app.Get("/tasks/:id", services.GetTask)
	//app.Patch("/updateTask", services.UpdateTask)
	//app.Delete("/deleteTask", services.DeleteTask)
	
}