package internal

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/ValeHenriquez/graph-gateway-server/config"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Response struct {
	Success string `json:"success"`
	Message string `json:"message"`
	Data    []byte `json:"data"`
}

func getQueueName(action_type string) (string, error) {
	queueTypes := map[string]string{
		"SIGNUP_USER":    "users_queue",
		"LOGIN_USER":     "users_queue",
		"GET_USER":       "users_queue",
		"GET_USERS":      "users_queue",
		"GET_USER_TASKS": "tasks_queue",
		"GET_TASK_USER":  "users_queue",
		"CREATE_TASK":    "tasks_queue",
		"GET_TASKS":      "tasks_queue",
		"GET_TASK":       "tasks_queue",
	}

	queueName, exists := queueTypes[action_type]
	if !exists {
		return "", errors.New("invalid action type")
	}

	return queueName, nil
}

func Handler(id string, body []byte, action_type string) (res Response, err error) {

	queue_name, err := getQueueName(action_type)

	if err != nil {
		return Response{}, err
	}

	if body == nil {
		var data struct {
			Id string `json:"id"`
		}
		data.Id = id
		body, err = json.Marshal(data)
		if err != nil {
			return Response{}, err
		}
	}

	ch := config.GetChannel()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	if err != nil {
		return Response{}, err
	}

	corrId := uuid.New().String()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(
		ctx,        // context
		"",         // exchange
		queue_name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "aplication/json",
			CorrelationId: corrId,
			ReplyTo:       q.Name,
			Body:          body,
			Type:          action_type,
		})

	if err != nil {
		return Response{}, err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return Response{}, err
	}

	for d := range msgs {
		if corrId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &res)
			break
		}
	}

	return res, nil
}
