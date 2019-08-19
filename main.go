// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

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

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "hello")
}

func main() {

	url := "mongodb://admin:kbr199sd5shi@localhost:27017"
	baseName := "data_test"
	mongoClient := mongodb.NewDBClient(url, baseName)
	mongoClient.PingTest()
	database := mongoClient.GetDatabase()
	collection := database.Collection("collection")

	connection := websocket.NewConnection("wss://api.huobi.br.com/ws")

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

	// 这里有个问题应该通过HeartBeat()，来判断当前状态

	router := httprouter.New()

	//针对subs
	router.GET("/", Index)
	router.GET("/subs", Hello)

	log.Fatal(http.ListenAndServe(":8080", router))
}
