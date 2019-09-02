package websocket

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"haki/common"
	"haki/mongodb"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

// WsConnection 所有的都是单例模式
var WsConnection *Connection

// Connection struct
type Connection struct {
	Ws *websocket.Conn

	ctx    context.Context
	cancel context.CancelFunc

	Addr  string
	Subs  []interface{}
	State int //反应状态
}

//Running 枚举
const (
	Running int = iota
	Stopped
)

// NewConnection return new connection
func NewConnection(url string) {
	ctx, cancel := context.WithCancel(context.TODO())
	WsConnection = &Connection{Addr: url, Ws: nil, State: Stopped, ctx: ctx, cancel: cancel}
}

// Connect will connect
func (c *Connection) Connect() error {

	wsConn, _, err := websocket.DefaultDialer.Dial(c.Addr, nil)

	if err != nil {
		fmt.Println(err)
		return err
	}
	c.Ws = wsConn
	c.State = Running
	return nil
}

// ReConnect will connect
func (c *Connection) ReConnect() error {
	fmt.Println("开始重连")
	wsConn, _, err := websocket.DefaultDialer.Dial(c.Addr, nil)
	if err != nil {
		return err
	}
	c.Ws = wsConn
	c.State = Running

	fmt.Println("重练完成")
	for _, sub := range c.Subs {
		log.Println("subscribe:", sub)
		c.Ws.WriteJSON(sub)
	}

	return nil
}

// CloseConnection will close c
func (c *Connection) CloseConnection() {
	c.cancel()
	err := c.Ws.Close()
	if err != nil {
		log.Println("close websocket connect error , ", err)
	}

	fmt.Println("关闭链接")
	c.State = Running
}

// HeartBeat 注册一个心跳函数用于检测链接，周期和
func (c *Connection) HeartBeat() {
	timer := time.NewTicker(time.Duration(5) * time.Second)

	// 心跳监控是对远程的
	go func() {
		fmt.Println("开启心跳监控")
		for {
			select {
			case <-timer.C:
				fmt.Println("检查远端状态")
				err := c.Ws.WriteJSON(map[string]interface{}{"ping": time.Duration(time.Now().Nanosecond())})
				if err != nil {
					fmt.Println("检查远端有问题")
					_ = c.ReConnect()
				}
			case <-c.ctx.Done():
				timer.Stop()
				log.Println("心跳监控关闭")
				return
			}
		}
	}()
}

// Subscribe will register subkey
func (c *Connection) Subscribe(subEvent interface{}) error {
	err := c.Ws.WriteJSON(subEvent)

	if err != nil {
		return err
	}

	c.Subs = append(c.Subs, subEvent)
	return nil
}

// Watch xxx
func (c *Connection) Watch() {
	go func() {
		for {
			select {
			case <-c.ctx.Done():
				return
			default:
				t, msg, err := c.Ws.ReadMessage()

				if err != nil {
					fmt.Println("重启")
					err := c.ReConnect()
					if err != nil {
						time.Sleep(1 * time.Second)
					}
					continue
				}
				switch t {
				case websocket.TextMessage, websocket.BinaryMessage:
					gzipreader, _ := gzip.NewReader(bytes.NewReader(msg))
					data, _ := ioutil.ReadAll(gzipreader)
					var resp map[string]interface{}
					json.Unmarshal(data, &resp)
					if resp["ping"] != nil {
						c.Ws.WriteJSON(map[string]interface{}{"pong": resp["ping"]})
					} else if resp["ch"] != nil {
						handleData(resp["ch"].(string), msg)
					} else if resp["pong"] != nil {
						c.State = Running
					} else {
						fmt.Println("错误:\n", resp)
					}

				case websocket.PongMessage:
					c.State = Running
				case websocket.CloseMessage:
					c.State = Stopped
				default:
					fmt.Println("错误:\n", string(msg))
				}
			}
		}
	}()
}

func handleData(ch string, msg []byte) {
	gzipreader, _ := gzip.NewReader(bytes.NewReader(msg))
	data, _ := ioutil.ReadAll(gzipreader)
	params := strings.Split(ch, ".")
	switch params[2] {
	case "kline":
		var kTicker common.KTicker
		json.Unmarshal(data, &kTicker)
		fmt.Println(kTicker)
		collection := mongodb.MongoClient.GetCollection("huobi", "kline")
		_, _ = collection.InsertOne(context.TODO(), kTicker)

	case "depth":
		var dTicker common.DTicker
		json.Unmarshal(data, &dTicker)
		fmt.Println(dTicker)
		collection := mongodb.MongoClient.GetCollection("huobi", "depth")
		_, _ = collection.InsertOne(context.TODO(), dTicker)
	case "trade":
		var tTicker common.TTicker
		json.Unmarshal(data, &tTicker)
		fmt.Println(tTicker)
		collection := mongodb.MongoClient.GetCollection("huobi", "trade")
		_, _ = collection.InsertOne(context.TODO(), tTicker)
	case "detail":
		var deTicker common.DeTicker
		json.Unmarshal(data, &deTicker)
		collection := mongodb.MongoClient.GetCollection("huobi", "detail")
		_, _ = collection.InsertOne(context.TODO(), deTicker)
	}
}
