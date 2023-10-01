package config

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

var ch *amqp.Channel
var conn *amqp.Connection

func SetupRabbitMQ() {
	URL := os.Getenv("RABBITMQ_URL")
	if URL == "" {
		failOnError(nil, "Failed to get RabbitMQ URL")
	}

	//fmt.Println(URL)
	conn, err := amqp.Dial(URL) //Establecer una conexion con el servidor de rabbitmq
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err = conn.Channel() //Sesion o instancia de comunicacion con el servidor de rabbitmq
	failOnError(err, "Failed to open a channel")

}

func GetChannel() *amqp.Channel {
	return ch
}

func CloseRabbitMQ() {
	ch.Close()
	conn.Close()
}
