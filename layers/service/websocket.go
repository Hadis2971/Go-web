package service

import "golang.org/x/net/websocket"

type Message struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type WebsocketService struct {
	Clients       map[string]map[*websocket.Conn]bool
	BroadcastChan chan Message
	ErrorChan chan bool
}

func NewWebsocketService() *WebsocketService {
	return &WebsocketService{
		Clients:       make(map[string]map[*websocket.Conn]bool),
		BroadcastChan: make(chan Message),
		ErrorChan: make(chan bool),
	}
}
