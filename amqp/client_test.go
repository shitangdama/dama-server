package amqp

import (
	"haki/config"
	"testing"
	"time"
)

type TransactionFinishedEvent struct {
	Id string `json:"id"`
}

func NewTransactionFinishedEvent() TransactionFinishedEvent {
	return TransactionFinishedEvent{
		Id: "1111",
	}
}

func TestNewClient(t *testing.T) {

	client := NewClient(config.RABBITMQ_URL, config.RABBITMQ_EXCHANGE)
	timer := time.NewTicker(time.Duration(1) * time.Second)

	for {
		select {
		case <-timer.C:
			payload := NewTransactionFinishedEvent()
			client.Publish("market.btcusdt.kline.1min", payload)
		}
	}
}
