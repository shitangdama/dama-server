package websocket

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

//

var handledata map[string]func(msg []byte)

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
func NewConnection(url string) *Connection {
	ctx, cancel := context.WithCancel(context.TODO())
	return &Connection{Addr: url, Ws: nil, State: Stopped, ctx: ctx, cancel: cancel}
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
func (c *Connection) Watch(handle func(msg []byte)) {
	go func() {
		for {
			select {
			case <-c.ctx.Done():
				return
			default:
				// err := c.Ws.SetReadDeadline(time.Now().Add(5 * time.Second))
				// if err != nil {
				// 	// 	c.ReConnect()
				// 	fmt.Println("重启2")
				// 	continue
				// }
				c.ReceiveMessage(handle)
			}
		}
	}()
}

// ReceiveMessage will handledata
func (c *Connection) ReceiveMessage(handle func(msg []byte)) {

	t, msg, err := c.Ws.ReadMessage()
	if err != nil {
		fmt.Println("重启")
		err := c.ReConnect()
		if err != nil {
			time.Sleep(1 * time.Second)
		}
		return
	}

	switch t {
	case websocket.TextMessage, websocket.BinaryMessage:
		handle(msg)
	case websocket.PongMessage:

	case websocket.CloseMessage:
		c.State = Stopped
	default:
		fmt.Println("错误:\n", string(msg))
	}
}
