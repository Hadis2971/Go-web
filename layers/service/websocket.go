package service

import "golang.org/x/net/websocket"

type Message struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type WebsocketService struct {
	Clients       map[string][]*websocket.Conn // Would be easier to have map[string]map[*websockeet.Conn]bool that way you don't have to loop through all connections when trying to find 1 of them
	BroadcastChan chan Message
}

func NewWebsocketService() *WebsocketService {
	return &WebsocketService{
		Clients:       make(map[string][]*websocket.Conn),
		BroadcastChan: make(chan Message),
	}
}
