package service

import (
	"github.com/Hadis2971/go_web/models"
	"golang.org/x/net/websocket"
)


type ChatMessage struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type ProductMessage struct {
	ID string `json:"id"`
	Topic string `json:"topic"`
	Product models.Product `json:"product"`
}

type Message interface {
	ChatMessage | ProductMessage
}

type WebsocketService[T Message] struct {
	Clients       map[string]map[*websocket.Conn]bool
	BroadcastChan chan T
	ErrorChan chan bool
}


func NewWebsocketService[T Message]() *WebsocketService[T] {
	return &WebsocketService[T]{
		Clients:       make(map[string]map[*websocket.Conn]bool),
		BroadcastChan: make(chan T),
		ErrorChan: make(chan bool),
	}
}
