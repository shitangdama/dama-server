package main

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"haki/common"
	"haki/mongodb"
	"haki/websocket"
)

var collection *websocket.Connection

// Index xx
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

// SubsIndex xx
func SubsIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "hello")
}

// SubsCreate xx
func SubsCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "hello")
}

// SubsDelete xx
func SubsDelete(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "hello")
}

func main() {

	client := NewClient(config.RABBITMQ_URL, config.RABBITMQ_EXCHANGE)

	defer client.Close()

	mongoClient := mongodb.NewDBClient(config.MONGODB_URL, config.MONGODB_DATABASE)
	database := mongoClient.GetDatabase()
	collection := database.Collection("collection")

	defer collection.CloseDatabase()

	connection := websocket.NewConnection(config.WS_COIN_URL)
	connection.Connect()
	connection.HeartBeat()

	defer connection.CloseConnection()

	connection.Subscribe(map[string]string{
		"sub": "market.htusdt.kline.1min",
		"id":  "1"})

	connection.Subscribe(map[string]string{
		"sub": "market.btcusdt.kline.1min",
		"id":  "1"})

	connection.Watch(func(msg []byte) {
		gzipreader, _ := gzip.NewReader(bytes.NewReader(msg))
		data, _ := ioutil.ReadAll(gzipreader)
		var resp map[string]interface{}
		json.Unmarshal(data, &resp)

		if resp["ping"] != nil {
			connection.Ws.WriteJSON(map[string]interface{}{"pong": resp["ping"]})
		} else if resp["ch"] != nil {
			// kv := strings.Split(resp["ch"].(string), ".")
			// fmt.Println(kv)
			var ticker common.Ticker
			json.Unmarshal(data, &ticker)
			fmt.Println(ticker)
			_, _ = collection.InsertOne(context.TODO(), ticker)
		} else {
		}
	})

	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/subs", SubsIndex)
	router.POST("/subs", SubsCreate)
	router.DELETE("/subs/:id", SubsDelete)

	log.Fatal(http.ListenAndServe(":8080", router))
}
