package service

import (
	"github.com/Hadis2971/go_web/models"
	"golang.org/x/net/websocket"
)

type ProductWsMessage struct {
	ID string `json:"id"`
	Topic string `json:"topic"`
	Product models.Product `json:"product"`
}

type ProductWebsocketService struct {
	Clients       map[string]map[*websocket.Conn]bool
	BroadcastChan chan ProductWsMessage
	ErrorChan chan bool
}

func NewProductWebsocketService() *ProductWebsocketService {
	return &ProductWebsocketService{
		Clients:       make(map[string]map[*websocket.Conn]bool),
		BroadcastChan: make(chan ProductWsMessage),
		ErrorChan: make(chan bool),
	}
}


