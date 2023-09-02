package routes

import (
	"github.com/ValeHenriquez/go101/ayu2/controllers"
	"github.com/ValeHenriquez/go101/ayu2/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	//Home route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	//Token validation Route (middleware example) 
	app.Get("/validate", middleware.Validate)

	//auth routes
	app.Post("/login", controllers.Login)
	app.Post("/signup", controllers.SingUp)

	//user routes
	app.Get("/users", controllers.GetUsers)
	app.Get("/users/:userId", controllers.GetUser)

	//task routes
	app.Get("/tasks", controllers.GetTasks)
	app.Get("/tasks/:taskId", controllers.GetTask)
	app.Post("/tasks", controllers.CreateTask)
	app.Put("/tasks/:taskId", controllers.UpdateTask)
	app.Delete("/tasks/:taskId", controllers.DeleteTask)

}
