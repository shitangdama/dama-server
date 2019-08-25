package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"haki/common"
	"haki/config"
	"haki/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

var connection *websocket.Connection

// NewRouter xx
func NewRouter() *httprouter.Router {
	router := httprouter.New()

	// for _, route := range routes {

	/*
		var handler mux.Handle

		handler = Logger(router, route.Method, route.Pattern, route.Name)
	*/
	// router.Handle(route.Method, route.Pattern, route.Handle)
	// }
	return router
}

func main() {

	// client := amqp.NewClient(config.RABBITMQ_URL, config.RABBITMQ_EXCHANGE)

	// defer client.Close()

	// mongoClient := mongodb.NewDBClient(config.MONGODB_URL, config.MONGODB_DATABASE)
	// database := mongoClient.GetDatabase()
	// collection := database.Collection("collection")

	// defer mongoClient.CloseDatabase()

	connection = websocket.NewConnection(config.WS_COIN_URL)
	connection.Connect()
	connection.HeartBeat()

	defer connection.CloseConnection()

	// connection.Subscribe(map[string]string{
	// 	"sub": "market.btcusdt.kline.1min",
	// 	"id":  "1"})

	// connection.Subscribe(map[string]string{
	// 	"sub": "market.btcusdt.depth.step0",
	// 	"id":  "1"})
	connection.Subscribe(map[string]string{
		"sub": "market.btcusdt.trade.detail",
		"id":  "1"})

	// connection.Subscribe(map[string]string{
	// 	"sub": "market.btcusdt.detail",
	// 	"id":  "1"})

	connection.Watch(func(msg []byte) {
		gzipreader, _ := gzip.NewReader(bytes.NewReader(msg))
		data, _ := ioutil.ReadAll(gzipreader)
		var resp map[string]interface{}
		json.Unmarshal(data, &resp)

		if resp["ping"] != nil {
			connection.Ws.WriteJSON(map[string]interface{}{"pong": resp["ping"]})
		} else if resp["ch"] != nil {

			params := strings.Split(resp["ch"].(string), ".")

			// kv := strings.Split(resp["ch"].(string), ".")
			// fmt.Println(kv)

			fmt.Println(params[2])
			switch params[2] {
			case "kline":
				var kTicker common.KTicker
				json.Unmarshal(data, &kTicker)
			case "depth":
				var dTicker common.DTicker
				json.Unmarshal(data, &dTicker)
				fmt.Println(dTicker)
			case "trade":
				var tTicker common.TTicker
				json.Unmarshal(data, &tTicker)
				fmt.Println(tTicker)
			case "detail":
				var deTicker common.DeTicker
				json.Unmarshal(data, &deTicker)
				// fmt.Println(dTicker)
			}

			// _, _ = collection.InsertOne(context.TODO(), ticker)
		} else {
			// fmt.Println(data)
		}
		fmt.Println(resp)
	})

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
