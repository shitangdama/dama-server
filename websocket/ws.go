package websocket

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

// Connection struct
type Connection struct {
	*websocket.Conn
	url     string
	subs    []interface{}
	close   chan int
	Actived time.Time
	isClose bool
}

// Endpoint 行情的Websocket入口
// var Endpoint = "wss://api.huobi.pro/ws"

// NewConnection return new Connection
func NewConnection(url string) *Connection {
	wsConn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return &Connection{url: url, Conn: wsConn}
}

// ReConnect 重新链接
func (c *Connection) ReConnect() error {
	wsConn, _, err := websocket.DefaultDialer.Dial(c.url, nil)
	c.Conn = wsConn
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Subscribe subscribe some topic
func (c *Connection) Subscribe(subEvent interface{}) error {
	err := c.WriteJSON(subEvent)
	if err != nil {
		return err
	}
	// c.subs = append(c.subs, subEvent)
	return nil
}

func (c *Connection) DisConnect() error {
	return nil
}

// 可以设计一个ctx，然后退出

// ReceiveMessage handle message from server
func (c *Connection) ReceiveMessage(handle func(msg []byte)) {
	for {
		t, msg, err := c.ReadMessage()

		if err != nil {
			log.Println(err)
			// if connection.isClose {
			// 	log.Println("exiting receive message goroutine.")
			// 				break
			// 			}
			// 			time.Sleep(time.Second)
			// 			continue
		}

		switch t {
		case websocket.TextMessage, websocket.BinaryMessage:
			handle(msg)
		// case websocket.PongMessage:
		// c.Actived = time.Now()
		case websocket.CloseMessage:
			// c.CloseWs()
			return
		default:
			log.Println("error websocket message type , content is :\n", string(msg))
		}
	}
}
