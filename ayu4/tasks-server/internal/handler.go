package internal

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/ValeHenriquez/example-rabbit-go/tasks-server/controllers"
	"github.com/ValeHenriquez/example-rabbit-go/tasks-server/models"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func Handler(d amqp.Delivery, ch *amqp.Channel) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var response models.Response

	actionType := d.Type
	switch actionType {
	case "CREATE_TASK":
		log.Println(" [.] Creating task")

		var task models.Task
		err := json.Unmarshal(d.Body, &task)
		failOnError(err, "Failed to Unmarshal task")
		taskJSON, err := json.Marshal(task)
		failOnError(err, "Failed to marshal task")

		_, err = controllers.CreateTask(task)
		if err != nil {
			response = models.Response{
				Success: "error",
				Message: "Task not created",
				Data:    []byte(err.Error()),
			}
		} else {
			response = models.Response{
				Success: "success",
				Message: "Task created",
				Data:    taskJSON,
			}
		}

	case "GET_TASKS":
		log.Println(" [.] Getting tasks")

		tasks, err := controllers.GetTasks()
		failOnError(err, "Failed to get tasks")
		tasksJSON, err := json.Marshal(tasks)
		failOnError(err, "Failed to marshal tasks")

		response = models.Response{
			Success: "success",
			Message: "Tasks retrieved",
			Data:    tasksJSON,
		}
	case "GET_TASK":
		log.Println(" [.] Getting task")
		var data struct {
			Id string `json:"id"`
		}
		var err error
		var taskJSON []byte
		var task models.Task

		err = json.Unmarshal(d.Body, &data)
		task, err = controllers.GetTask(data.Id)
		

		taskJSON, err = json.Marshal(task)
		if err != nil {
			response = models.Response{
				Success: "error",
				Message: "Task not created",
				Data:    []byte(err.Error()),
			}
		} else {
			response = models.Response{
				Success: "success",
				Message: "Task retrieved",
				Data:    taskJSON,
			}
		}

	case "GET_USER_TASKS":
		log.Println(" [.] Getting user tasks")
		var requestData struct {
			ID string `json:"id"`
		}

		err := json.Unmarshal([]byte(d.Body), &requestData)
		failOnError(err, "Failed to Unmarshal task")
		tasks, err := controllers.GetUserTasks(requestData.ID)
		failOnError(err, "Failed to get tasks")
		tasksJSON, err := json.Marshal(tasks)
		failOnError(err, "Failed to marshal tasks")
		response = models.Response{
			Success: "success",
			Message: "Tasks retrieved",
			Data:    tasksJSON,
		}
	}

	responseJSON, err := json.Marshal(response)
	failOnError(err, "Failed to marshal response")

	err = ch.PublishWithContext(ctx,
		"",        // exchange
		d.ReplyTo, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: d.CorrelationId,
			Body:          responseJSON,
		})
	failOnError(err, "Failed to publish a message")

	d.Ack(false)
}
