package domain

import (
	"sync"

	"golang.org/x/net/websocket"

	"github.com/Hadis2971/go_web/layers/service"
)

type ChatDomain struct {
	websocket *service.WebsocketService
	mutex     sync.Mutex
}

func NewChatDomain(websocket *service.WebsocketService) *ChatDomain {
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

	
	
	// There's no authentication for this, so you can end up with infinite connections.
	// Hadis => How would I do the auth? How can I comapre 2 ws conns?

	// if cd.websocket.Clients[id][conn] {
	// 	return
	// }
	// Can I do something like this?
	// It seems that golang does a deep comparison instead of a shallow. So it would work?

	cd.mutex.Unlock()

	go cd.start() 
	
	// This doesn't look correct, you have 2 golang loops, and when does this loop stop? Never?
	// Hadis => I added a new chan for errors so I add a true values on line 37 when there is an error
	// 			So if the client shutsdown the connection add to the chan true and then i have a case in
	//			start for it and do a return

	for {
		var message service.Message

		if err := websocket.JSON.Receive(conn, &message); err != nil {
			conn.Close()
			cd.removeClient(message.ID, conn)
	
			cd.websocket.ErrorChan <- true

			return
		}

		cd.websocket.BroadcastChan <- message
	}
}

func (cd *ChatDomain) removeClient(id string, conn *websocket.Conn) {
	cd.mutex.Lock()

	clients := cd.websocket.Clients[id]

	delete(clients, conn)

	cd.websocket.Clients[id] = clients

	cd.mutex.Unlock()
}

func (cd *ChatDomain) start() {
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

func (cd *ChatDomain) handleBroadcastMsg(msg service.Message) {
	clients := cd.websocket.Clients[msg.ID]

	for client := range clients {
		websocket.JSON.Send(client, msg.Text)
	}
}
