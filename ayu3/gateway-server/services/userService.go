package services

import (
	"github.com/ValeHenriquez/example-rabbit-go/gateway-server/internal"
	"github.com/gofiber/fiber/v2"
)


func Login(c *fiber.Ctx) error {
	actionType := "LOGIN_USER"
	return internal.HandleRPCRequest(0, c, actionType)
}

func SignUp(c *fiber.Ctx) error {
	actionType := "SIGNUP_USER"
	return internal.HandleRPCRequest(0, c, actionType)
}

