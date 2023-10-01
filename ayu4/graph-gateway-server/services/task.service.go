package services

import (
	"encoding/json"
	"log"

	"github.com/ValeHenriquez/graph-gateway-server/graph/model"
	"github.com/ValeHenriquez/graph-gateway-server/internal"
)

func HandleCreateTask(input model.NewTask) (*model.Task, error) {
	action_type := "CREATE_TASK"
	body, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	response, err := internal.Handler("", body, action_type)
	if err != nil {
		return nil, err
	}
	var task *model.Task
	err = json.Unmarshal(response.Data, &task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func HandleTasks() ([]*model.Task, error) {
	action_type := "GET_TASKS"
	response, err := internal.Handler("", nil, action_type)
	if err != nil {
		return nil, err
	}
	var tasks []*model.Task
	err = json.Unmarshal(response.Data, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func HandleTask(id string) (*model.Task, error) {
	action_type := "GET_TASK"
	response, err := internal.Handler(id, nil, action_type)
	if err != nil {
		return nil, err
	}
	var task *model.Task
	err = json.Unmarshal(response.Data, &task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

// Tareas de un usuario
func HandleUserTasks(userID string) ([]*model.Task, error) {
	action_type := "GET_USER_TASKS"
	log.Println("userID: ", userID)
	response, err := internal.Handler(userID, nil, action_type)
	if err != nil {
		return nil, err
	}
	var tasks []*model.Task
	err = json.Unmarshal(response.Data, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
