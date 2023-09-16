package controllers

import (
	db "github.com/ValeHenriquez/example-rabbit-go/tasks-server/config"
	"github.com/ValeHenriquez/example-rabbit-go/tasks-server/models"
)

func GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := db.DB.Select("id, title, description, done, user_id").Find(&tasks).Error

	return tasks, err
}

func GetTask(taskID uint) (models.Task, error) {
	var task models.Task
	err := db.DB.Select("id, title, description, done").Where("id = ?", taskID).First(&task).Error

	return task, err
}

func CreateTask(task models.Task) (models.Task, error) {
	err := db.DB.Create(&task).Error

	return task, err
}

func UpdateTask(taskID uint, updatedTask models.Task) (models.Task, error) {
	var task models.Task
	err := db.DB.Select("id, title, description, completed").Where("id = ?", taskID).First(&task).Error

	if err != nil {
		return task, err
	}

	err = db.DB.Model(&task).Updates(&updatedTask).Error

	return task, err
}

func DeleteTask(taskID uint) error {
	var task models.Task
	err := db.DB.Select("id, title, description, completed").Where("id = ?", taskID).First(&task).Error

	if err != nil {
		return err
	}

	err = db.DB.Delete(&task).Error

	return err
}
