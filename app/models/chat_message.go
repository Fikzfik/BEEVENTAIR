package models

import "time"

type ChatMessage struct {
	ID          string    `json:"id"`
	ChannelName string    `json:"channel_name"`
	SenderID    string    `json:"sender_id"`
	Content     string    `json:"content"`
	Timestamp   time.Time `json:"timestamp"`
}
