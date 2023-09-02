package controllers

import (
	"net/http"
	"os"
	"time"

	db "github.com/ValeHenriquez/go101/ayu2/config"
	"github.com/ValeHenriquez/go101/ayu2/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	//Obtener el correo y contraseña desde el body del request

	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if c.BodyParser(&body) != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Bad request"})
	}

	//Buscar al usuario en la DB
	var user models.User
	db.DB.Where("email = ?", body.Email).First(&user)

	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "User not found"})
	}

	//Verificar la contraseña

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{"message": "Incorrect password"})

	}

	//Generar JWT Token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Issuer":    user.ID,
		"ExpiresAt": time.Now().Add(time.Hour * 24 * 30).Unix(), //30 days
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Token Expired or invalid",
		})
	}

	//Configurar y setear la cookie
	c.Cookie(&fiber.Cookie{
		Name:     "Authorization",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24 * 30),
	})

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Logged in",
		"data":    user,
	})

}

func SingUp(c *fiber.Ctx) error {
	//Obtener los campos desde el body del request

	var body struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	if c.BodyParser(&body) != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Bad request"})
	}

	// Generar un hash de la contraseña para almacenarla de manera segura
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Internal server error"})
	}

	// Crear el usuario en la DB
	user := models.User{
		Email:     body.Email,
		Password:  string(hash),
		FirstName: body.FirstName,
		LastName:  body.LastName,
	}
	result := db.DB.Create(&user)


	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Internal server error"})
	}
	
	//Respuesta exitosa
	return c.Status(http.StatusCreated).JSON(user)

}
