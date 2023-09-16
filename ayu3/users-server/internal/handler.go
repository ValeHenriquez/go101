package internal

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/ValeHenriquez/example-rabbit-go/users-server/controllers"
	"github.com/ValeHenriquez/example-rabbit-go/users-server/models"
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
	var tokenString string
	actionType := d.Type
	switch actionType {
	case "LOGIN_USER":
		log.Println(" [.] Login user")

		var err error
		tokenString, err = controllers.Login(d.Body)

		if err != nil {
			response = models.Response{
				Success: "error",
				Message: "Login failed",
				Data:    []byte(err.Error()),
			}
		} else {
			user, err := controllers.GetUserByEmail(d.Body)
			JSONUser, err := json.Marshal(user)
			failOnError(err, "Failed to marshal user")
			response = models.Response{
				Success: "success",
				Message: "User logged",
				Data:    JSONUser,
			}
		}

	case "SIGNUP_USER":
		log.Println(" [.] Creating User")

		err := controllers.SingUp(d.Body)
		failOnError(err, "Failed to create User")

		response = models.Response{
			Success: "success",
			Message: "User Created",
			Data:    nil,
		}
	case "GET_USER":
		var data struct {
			ID uint `json:"id"`
		}

		log.Println(" [.] Getting user")
		err := json.Unmarshal(d.Body, &data)
		failOnError(err, "Failed to Unmarshal user")

		user, err := controllers.GetUser(data.ID)
		failOnError(err, "Failed to get user")
		userJSON, err := json.Marshal(user)
		failOnError(err, "Failed to marshal user")

		response = models.Response{
			Success: "success",
			Message: "User retrieved",
			Data:    userJSON,
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
			Headers:       amqp.Table{"Authorization": tokenString},
		})
	failOnError(err, "Failed to publish a message")

	d.Ack(false)
}
