package services

import (
	"encoding/json"

	"github.com/ValeHenriquez/graph-gateway-server/graph/model"
	"github.com/ValeHenriquez/graph-gateway-server/internal"
)

func HandleSignUp(input model.NewUser) (*model.User, error) {
	action_type := "SIGNUP_USER"
	body, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	response, err := internal.Handler("", body, action_type)
	if err != nil {
		return nil, err
	}
	var user *model.User
	err = json.Unmarshal(response.Data, &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func HandleLogin(input model.LoginInput) (string, error) {
	action_type := "LOGIN_USER"
	body, err := json.Marshal(input)
	if err != nil {
		return "", err
	}
	response, err := internal.Handler("", body, action_type)
	if err != nil {
		return "", err
	}
	var token string
	err = json.Unmarshal(response.Data, &token)
	if err != nil {
		return "", err
	}
	return token, nil
}