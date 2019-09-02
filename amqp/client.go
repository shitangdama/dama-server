package amqp

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

// AmqpClient xx
var AmqpClient *amqpClient

type amqpClient struct {
	Connection *amqp.Connection
	Exchange   string
	Channel    *amqp.Channel
}

// NewClient xx amqpClient
func NewClient(url string, exchange string) {
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to declare a queue")
	err = ch.ExchangeDeclare(
		exchange, // name
		"topic",
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")
	AmqpClient = &amqpClient{Connection: conn, Exchange: exchange, Channel: ch}
	// return &amqpClient{Connection: conn, Exchange: exchange, Channel: ch}
}

func (client *amqpClient) Close() {
	client.Channel.Close()
	client.Connection.Close()
}

func (client *amqpClient) Publish(routingKey string, params interface{}) error {

	data, err := json.Marshal(params)
	if err != nil {
		// Failed to encode payload
		return err
	}

	publishing := amqp.Publishing{
		ContentType:     "application/json",
		ContentEncoding: "UTF-8",
		Body:            data,
	}

	return client.Channel.Publish(client.Exchange, routingKey, false, false, publishing)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
