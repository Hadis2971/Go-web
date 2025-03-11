package domain

import (
	"sync"

	"golang.org/x/net/websocket"

	"github.com/Hadis2971/go_web/layers/service"
)

type ChatDomain struct {
	websocket *service.WebsocketService
	mutex sync.Mutex
}

func NewChatDomain (websocket *service.WebsocketService) *ChatDomain {
	return &ChatDomain{websocket: websocket}
}

func (cd *ChatDomain) AddNewClient (id string , conn *websocket.Conn) {
	cd.mutex.Lock()

	cd.websocket.Clients[id] = append(cd.websocket.Clients[id], conn)

	cd.mutex.Unlock()

	go cd.start()

	for {
		var message service.Message

		if err := websocket.JSON.Receive(conn, &message); err != nil {
			cd.removeClient(message.ID, conn)

			return
		}

		cd.websocket.BroadcastChan <- message
	}
}

func (cd * ChatDomain) removeClient (id string, conn *websocket.Conn) {
	cd.mutex.Lock()

	clients := cd.websocket.Clients[id];

	var index int

	for idx, _ := range clients {
		if clients[idx] == conn {
			index = idx
			break
		}
	}

	clients = append(clients[:index], clients[:index + 1]...)
	cd.websocket.Clients[id] = clients;
	
	cd.mutex.Unlock()
}

func (cd *ChatDomain) start () {
	for {
		select {
		case message := <- cd.websocket.BroadcastChan:
			cd.handleBroadcastMsg(message)
		}
	}
}

func (cd *ChatDomain) handleBroadcastMsg (msg service.Message) {
	clients := cd.websocket.Clients[msg.ID]

	for _, client := range clients {
		websocket.JSON.Send(client, msg.Text)
	} 
}