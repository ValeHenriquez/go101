package main

import (
	"fmt"

	"github.com/ValeHenriquez/example-rabbit-go/gateway-server/config"
	"github.com/ValeHenriquez/example-rabbit-go/gateway-server/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)



func main() {
	fmt.Println("API Gateway starting...")
	godotenv.Load()
	fmt.Println("Loaded env variables...")

	fmt.Println("Configuring RabbitMQ Connection...")
	config.SetupRabbitMQ()
	fmt.Println("RabbitMQ Connection configured...")

	app := fiber.New()

	routes.Setup(app)

	app.Use(cors.New())
	app.Listen(":3000")
	fmt.Println("API Gateway started... Listening op port 3000")

	defer config.CloseRabbitMQ()
}

