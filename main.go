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
	"test/common"
	"test/mongodb"
	"test/websocket"
)

func main() {
	connection := websocket.NewConnection("wss://api.huobi.br.com/ws")
	url := "mongodb://admin:kbr199sd5shi@localhost:27017"

	baseName := "data_test"
	mongoClient := mongodb.NewDBClient(url, baseName)
	database := mongoClient.GetDatabase()
	collection := database.Collection("collection")

	// // connection.Subscribe(map[string]string{
	// // 	"sub": "market.btcusdt.detail",
	// // 	"id":  "2"})
	// // connection.Subscribe(map[string]string{
	// // 	"sub": "market.btcusdt.depth.step0",
	// // 	"id":  "1"})
	connection.Subscribe(map[string]string{
		"sub": "market.htusdt.kline.1min",
		"id":  "1"})

	connection.ReceiveMessage(func(msg []byte) {
		gzipreader, _ := gzip.NewReader(bytes.NewReader(msg))
		data, _ := ioutil.ReadAll(gzipreader)
		var resp map[string]interface{}
		json.Unmarshal(data, &resp)

		if resp["ping"] != nil {
			connection.WriteJSON(map[string]interface{}{"pong": resp["ping"]})
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

}
