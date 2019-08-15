// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	"haki/common"
	"haki/websocket"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "hello")
}

func main() {

	// router := httprouter.New()

	// // 主要是针对sub上面的
	// // 现只有查看sub对
	// router.GET("/", Index)
	// router.GET("/subs", Hello)

	// log.Fatal(http.ListenAndServe(":8080", router))

	connection := websocket.NewConnection("wss://api.huobi.br.com/ws")

	connection.Connect()

	connection.Subscribe(map[string]string{
		"sub": "market.htusdt.kline.1min",
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
			// _, _ = collection.InsertOne(context.TODO(), ticker)
		} else {
		}
	})

	// 这里有个问题应该通过HeartBeat()，来判断当前状态
	connection.HeartBeat()
	time.Sleep(5 * time.Second)
	// fmt.Println("数据断线")
	// connection.State = websocket.Stopped
	// fmt.Println("数据重启")
	time.Sleep(50 * time.Second)

	connection.CloseConnection()

}
