package websocket

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"haki/websocket"
	"io/ioutil"
	"test/common"
	"testing"
	"time"
)

func TestNewConnection(t *testing.T) {

	connection := websocket.NewConnection("wss://api.huobi.br.com/ws")

	connection.Connect()

	connection.Subscribe(map[string]string{
		"sub": "market.htusdt.kline.1min",
		"id":  "1"})

	connection.ReceiveMessage(func(msg []byte) {
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
			t.Log(ticker)
			// _, _ = collection.InsertOne(context.TODO(), ticker)
		} else {
		}
	})

	//测试中断
	time.Sleep(3 * time.Second)
	t.Log(11111)
	connection.CloseConnection()
	t.Log(11111)
	// time.Sleep(2 * time.Second)
}
