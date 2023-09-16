package controllers

import (
	"encoding/json"
	"errors"
	"os"
	"time"

	db "github.com/ValeHenriquez/example-rabbit-go/users-server/config"
	"github.com/ValeHenriquez/example-rabbit-go/users-server/models"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)





func Login(body []byte) (string, error) {
	//Obtener el correo y contraseña desde el body del request
	var loginData models.LoginData

	err := json.Unmarshal(body, &loginData); 
	if err != nil {
		return "",err
	}
	//Buscar al usuario en la DB
	var user models.User
	db.DB.Where("email = ?", loginData.Email).First(&user)

	if user.ID == 0 {
		return "",errors.New("User not found")
	}

	//Verificar la contraseña

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		return "",errors.New("Incorrect password")

	}

	//Generar JWT Token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Issuer":    user.ID,
		"ExpiresAt": time.Now().Add(time.Hour * 24 * 30).Unix(), //30 days
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "",errors.New("Token Expired or invalid")
	}

	return tokenString, nil

}

func SingUp(body []byte) error {
	//Obtener los campos desde el body del request

	var singUpData models.SingUpData

	err := json.Unmarshal(body, &singUpData); 
	if err != nil {
		return err
	}
	
	// Generar un hash de la contraseña para almacenarla de manera segura
	hash, err := bcrypt.GenerateFromPassword([]byte(singUpData.Password), bcrypt.DefaultCost)

	if err != nil {
		return errors.New("Failed to hash password")
	}

	// Crear el usuario en la DB
	user := models.User{
		Email:     singUpData.Email,
		Password:  string(hash),
		FirstName: singUpData.FirstName,
		LastName:  singUpData.LastName,
	}
	result := db.DB.Create(&user)


	if result.Error != nil {
		return errors.New("Failed to create user")
	}
	
	//Respuesta exitosa
	return nil

}
