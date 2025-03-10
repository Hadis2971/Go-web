package service

import "golang.org/x/net/websocket"

type Message struct {
	ID string
	Text string
}

type WebsocketService struct {
	Clients map[string][]*websocket.Conn
	// addClientChan chan *websocket.Conn
	// removeClientChan chan *websocket.Conn
	BroadcastChan chan Message
}

func NewWebsocketService () *WebsocketService {
	return &WebsocketService{
		Clients: make(map[string][]*websocket.Conn),
		// addClientChan: make(chan *websocket.Conn),
		// removeClientChan: make(chan *websocket.Conn),
		BroadcastChan: make(chan Message),
	}
}