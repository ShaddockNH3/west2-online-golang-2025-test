package db

import "time"

type WebSocketMessage struct {
	ID          string // UUID
	MessageType int64  // 目前只支持0-5
	SenderID    string
	TargetID    string // 目标用户ID或群ID
	Content     string
	IsRead      bool // 是否已读
	CreateAt    time.Time
}

type Room struct {
	ID        string // UUID
	Name      string
	Members   []string // 成员用户ID列表
	CreateAt  time.Time
}

// 创建房间

// 保存消息

// 查询历史记录
