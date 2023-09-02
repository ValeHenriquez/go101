package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)


func AuthenticateToken(tokenString string) error {
	//token check
	_, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return err
	}
	return nil

}

func Validate(c *fiber.Ctx) error {
	//get token from cookies
	token := c.Cookies("Authorization")
	if token == "" {
		return c.Status(401).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized",
		})
	}
	err := AuthenticateToken(token)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized",
		})
	} else {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Authorized",
		})
	}

}
