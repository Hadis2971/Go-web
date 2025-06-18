package domain

import (
	"sync"

	"golang.org/x/net/websocket"

	"github.com/Hadis2971/go_web/layers/service"
)

type ChatDomain struct {
	websocket *service.WebsocketService[service.ChatMessage]
	mutex     sync.Mutex
}

func NewChatDomain(websocket *service.WebsocketService[service.ChatMessage]) *ChatDomain {
	return &ChatDomain{websocket: websocket}
}

func (cd *ChatDomain) AddNewClient(id string, conn *websocket.Conn) {
	cd.mutex.Lock()

	if (cd.websocket.Clients[id] == nil) {
		cd.websocket.Clients[id] = make(map[*websocket.Conn]bool)
	}

	if (cd.websocket.Clients[id][conn] == true) {
		return
	} else {
		cd.websocket.Clients[id][conn] = true
	}

	cd.mutex.Unlock()

	go cd.startBroadcaseChatWsMsgs() 
	go cd.startReceiveChatWsMsgs(conn)
	
}

func (cd *ChatDomain) removeClient(id string, conn *websocket.Conn) {
	cd.mutex.Lock()

	delete(cd.websocket.Clients[id], conn)

	if (len(cd.websocket.Clients[id]) == 0) {
		delete(cd.websocket.Clients, id)
	}

	cd.mutex.Unlock()
}

func (cd *ChatDomain) startBroadcaseChatWsMsgs() {
	for {
		select {
		case message := <-cd.websocket.BroadcastChan:
			cd.handleBroadcastMsg(message)
		case err := <- cd.websocket.ErrorChan:
			if err {
				return
			}
		}
	}
}

func (cd *ChatDomain) startReceiveChatWsMsgs(conn *websocket.Conn) {
	for {
		var message service.ChatMessage

		if err := websocket.JSON.Receive(conn, &message); err != nil {
			conn.Close()
			cd.removeClient(message.ID, conn)
	
			cd.websocket.ErrorChan <- true

			return
		}

		cd.websocket.BroadcastChan <- message
	}
}

func (cd *ChatDomain) handleBroadcastMsg(msg service.ChatMessage) {
	clients := cd.websocket.Clients[msg.ID]

	for client := range clients {
		websocket.JSON.Send(client, msg.Text)
	}
}
