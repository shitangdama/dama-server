package amqp

import (
	"haki/config"
	"testing"
	"time"
)

type TransactionFinishedEvent struct {
	ID string `json:"id"`
}

func NewTransactionFinishedEvent() TransactionFinishedEvent {
	return TransactionFinishedEvent{
		ID: "1111",
	}
}

func TestNewClient(t *testing.T) {

	NewClient(config.RABBITMQ_URL, config.RABBITMQ_EXCHANGE)
	timer := time.NewTicker(time.Duration(1) * time.Second)

	for {
		select {
		case <-timer.C:
			payload := NewTransactionFinishedEvent()
			AmqpClient.Publish("market.btcusdt.kline.1min", payload)
		}
	}
}
