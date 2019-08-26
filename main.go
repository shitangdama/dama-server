package main

import (
	"haki/config"
	"haki/mongodb"
	"haki/router"
	"haki/websocket"
	"log"
	"net/http"
)

func main() {

	// client := amqp.NewClient(config.RABBITMQ_URL, config.RABBITMQ_EXCHANGE)

	// defer client.Close()

	mongodb.NewDBClient(config.MONGODB_URL, config.MONGODB_DATABASE)
	// database := mongoClient.GetDatabase()
	// collection := database.Collection("collection")

	// defer mongoClient.CloseDatabase()

	websocket.NewConnection(config.WS_COIN_URL)
	websocket.WsConnection.Connect()
	websocket.WsConnection.HeartBeat()

	defer websocket.WsConnection.CloseConnection()

	websocket.WsConnection.Subscribe(map[string]string{
		"sub": "market.btcusdt.kline.1min",
		"id":  "1"})

	websocket.WsConnection.Subscribe(map[string]string{
		"sub": "market.btcusdt.depth.step0",
		"id":  "1"})

	websocket.WsConnection.Subscribe(map[string]string{
		"sub": "market.btcusdt.trade.detail",
		"id":  "1"})

	websocket.WsConnection.Subscribe(map[string]string{
		"sub": "market.btcusdt.detail",
		"id":  "1"})

	websocket.WsConnection.Watch()

	router := router.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
