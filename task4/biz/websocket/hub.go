package websocket

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/biz/dal/db"
)

type Hub struct {
	// 注册的客户端，通过uuid识别
	clients     map[string]*Client
	clientsLock sync.RWMutex

	// 群聊客户端记录了每个群聊里都有哪些成员
	rooms     map[string]map[string]struct{}
	roomsLock sync.RWMutex

	// 广播消息频道
	broadcast chan *Message

	register   chan *Client
	unregister chan *Client
}

type Message struct {
	Data     []byte // 原始消息数据
	SenderID string // 发送者的 UserID
}

// ClientMessage DTO (Data Transfer Object)
type ClientMessage struct {
	MessageType int64  `json:"message_type"`
	TargetID    string `json:"target_id"`
	Content     string `json:"content"`
}

var hub = NewHub()

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan *Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[string]*Client),
		rooms:      make(map[string]map[string]struct{}),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.Register(client)
		case client := <-h.unregister:
			h.Unregister(client)
		case message := <-h.broadcast:
			var clientMessage ClientMessage
			json.Unmarshal(message.Data, &clientMessage)
			// 存入数据库的逻辑
			fullMessage := &db.WebSocketMessage{
				ID:          "", // 这里可以生成一个 UUID
				MessageType: clientMessage.MessageType,
				SenderID:    message.SenderID,
				TargetID:    clientMessage.TargetID,
				Content:     clientMessage.Content,
				IsRead:      false,
				CreateAt:    time.Now(),
			}

			// 存入数据库的逻辑暂时省略
			_ = fullMessage // 这里可以调用数据库保存函数

			switch clientMessage.MessageType {
			case 0:
				// 处理类型0的消息
				for _, client := range h.clients {
					select {
					case client.send <- message.Data:
					default:
						close(client.send)
						delete(h.clients, client.UserID)
					}
				}
			case 1:
				// 处理类型1的消息

			case 2:
				// 处理类型2的消息
			case 3:
				// 处理类型3的消息
			case 4:
				// 处理类型4的消息
			case 5:
				// 处理类型5的消息
			default:
				// 未知消息类型，记录错误日志
			}
		}
	}
}

func (h *Hub) Register(client *Client) {
	h.AddClient(client)
}

func (h *Hub) AddClient(client *Client) {
	// 添加客户端到客户端集合，并发安全，使用读写锁保护
	h.clientsLock.Lock()
	defer h.clientsLock.Unlock()
	h.clients[client.UserID] = client
}

func (h *Hub) Unregister(client *Client) {
	h.DelClient(client)
}

func (h *Hub) DelClient(client *Client) {
	// 从客户端集合中删除客户端，并发安全，使用读写锁保护
	h.clientsLock.Lock()
	defer h.clientsLock.Unlock()
	if _, ok := h.clients[client.UserID]; ok {
		delete(h.clients, client.UserID)
		close(client.send)
	}
}
