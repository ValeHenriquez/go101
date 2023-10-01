package services

import (
	"encoding/json"

	"github.com/ValeHenriquez/graph-gateway-server/graph/model"
	"github.com/ValeHenriquez/graph-gateway-server/internal"
)

func HandleUsers () ([]*model.User, error) {
	action_type := "GET_USERS"
	response, err := internal.Handler("", nil, action_type)
	if err != nil {
		return nil, err
	}
	var users []*model.User
	err = json.Unmarshal(response.Data, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func HandleUser (id string) (*model.User, error) {
	action_type := "GET_USER"
	response, err := internal.Handler(id, nil, action_type)
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

