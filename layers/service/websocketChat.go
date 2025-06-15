package service

import "golang.org/x/net/websocket"

type ChatWsMessage struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type WebsocketChatService struct {
	Clients       map[string]map[*websocket.Conn]bool
	BroadcastChan chan ChatWsMessage
	ErrorChan chan bool
}

func NewWebsocketService() *WebsocketChatService {
	return &WebsocketChatService{
		Clients:       make(map[string]map[*websocket.Conn]bool),
		BroadcastChan: make(chan ChatWsMessage),
		ErrorChan: make(chan bool),
	}
}
