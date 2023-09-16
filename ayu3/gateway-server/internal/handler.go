package internal

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/ValeHenriquez/example-rabbit-go/gateway-server/config"
	"github.com/ValeHenriquez/example-rabbit-go/gateway-server/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)



func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}


func getQueueName(action_type string) (string, error) {
	queueTypes := map[string]string{
		"SIGNUP_USER": "users_queue",
		"LOGIN_USER":  "users_queue",
		"CREATE_TASK": "tasks_queue",
		"GET_TASKS":   "tasks_queue",
		"GET_TASK":    "tasks_queue",
	}

	queueName, exists := queueTypes[action_type]
	if !exists {
		return "", errors.New("invalid action type")
	}

	return queueName, nil
}


func handler(id uint64, body []byte, action_type string) (res models.Response, err error) {

	queue_name,err := getQueueName(action_type);
	failOnError(err, "Failed to get queue name")

	ch := config.GetChannel()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	corrId := uuid.New().String()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	
	err = ch.PublishWithContext(
		ctx,   // context
		"",    // exchange
		queue_name, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType:   "aplication/json",
			CorrelationId: corrId,
			ReplyTo:       q.Name,
			Body:          body,
			Type:          action_type,
		})

	failOnError(err, "Failed to publish a message")
	

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	for d := range msgs {
		if corrId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &res)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return
}


func HandleRPCRequest(id uint64, c *fiber.Ctx, actionType string) error {
	body := c.Body()
	res, err := handler(id, body, actionType)
	if err != nil {
		log.Printf("Failed to handle RPC request: %s", err)
		return c.Status(500).JSON(models.Response{
			Success: "false",
			Message: "Internal Server Error",
		})
	}
	return c.Status(200).JSON(res)
}