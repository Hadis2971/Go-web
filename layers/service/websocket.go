package service

import "golang.org/x/net/websocket"

type Message struct {
	ID string `json:"id"`
	Text string `json:"text"`
}

type WebsocketService struct {
	Clients map[string][]*websocket.Conn
	BroadcastChan chan Message
}

func NewWebsocketService () *WebsocketService {
	return &WebsocketService{
		Clients: make(map[string][]*websocket.Conn),
		BroadcastChan: make(chan Message),
	}
}