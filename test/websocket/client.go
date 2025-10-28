package websocket

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/pkg/constants"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/websocket"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// Client 表示一个 WebSocket 客户端连接
type Client struct {
	conn   *websocket.Conn
	send   chan []byte
	UserID string // 连接对应的用户ID
}

// ReadPump 读取来自 WebSocket 连接的消息，并将其发送到 Hub 的广播频道
func (c *Client) ReadPump() {
	defer func() {
		hub.unregister <- c
		c.conn.Close()
	}()
	// 设置读取限制和超时，以及 Pong 处理程序，以支持心跳机制
	c.conn.SetReadLimit(constants.MaxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(constants.PongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(constants.PongWait))
		return nil
	})
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				// 记录错误日志
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		var msg Message
		err = json.Unmarshal(message, &msg)
		hub.broadcast <- &msg
	}
}

// WritePump 处理来自 Hub 的消息，并将其写入 WebSocket 连接，包括心跳机制
func (c *Client) WritePump() {
	ticker := time.NewTicker(constants.PingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(constants.WriteWait))
			if !ok {
				// Hub 关闭了频道，关闭连接
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// 添加队列中的消息
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			// 关闭写入器
			if err := w.Close(); err != nil {
				return
			}

		// 发送心跳，保持连接活跃，防止超时断开
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(constants.WriteWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func serveWs(ctx *app.RequestContext) {
	err := upgrader.Upgrade(ctx, func(conn *websocket.Conn) {
		client := &Client{conn: conn, send: make(chan []byte, 256)}
		hub.register <- client

		go client.WritePump()
		client.ReadPump()
	})
	if err != nil {
		// 记录错误日志
		return
	}
}
