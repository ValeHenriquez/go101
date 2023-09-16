package controllers

import (
	"encoding/json"

	db "github.com/ValeHenriquez/example-rabbit-go/users-server/config"
	"github.com/ValeHenriquez/example-rabbit-go/users-server/models"
)

func GetUsers() ([]models.User, error) {
	var users []models.User
	err := db.DB.Select("id, first_name, last_name, email").Find(&users).Error
	return users, err
}

func GetUser(userkID uint) (models.User, error) {
	var user models.User
	err := db.DB.Select("id, title, description, done").Where("id = ?", userkID).First(&user).Error

	return user, err
}

func GetUserByEmail(body []byte) (models.User, error){		
	var loginData models.LoginData

	err := json.Unmarshal(body, &loginData); 
	if err != nil {
		return models.User{},err
	}

	var user models.User
	err = db.DB.Select("id, first_name, last_name, email").Where("email = ?", loginData.Email).First(&user).Error
	return user, err
}