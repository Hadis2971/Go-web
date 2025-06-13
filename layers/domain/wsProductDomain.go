package domain

import (
	"fmt"
	"sync"

	"golang.org/x/net/websocket"

	"github.com/Hadis2971/go_web/layers/service"
)

type WsProductDomain struct {
	websocket *service.ProductWebsocketService
	mutex     sync.Mutex
}

func NewWsProductDomain(websocket *service.ProductWebsocketService) *WsProductDomain {
	return &WsProductDomain{websocket: websocket}
}

func (wspd *WsProductDomain) AddNewWsProductClient(id string, conn *websocket.Conn) {

	fmt.Println("HAHAH ID", id)

	wspd.mutex.Lock()

	if (wspd.websocket.Clients[id] == nil) {
		wspd.websocket.Clients[id] = make(map[*websocket.Conn]bool)
	}

	if (wspd.websocket.Clients[id][conn] == true) {
		return
	} else {
		wspd.websocket.Clients[id][conn] = true
	}

	fmt.Println("wspd.websocket.Clients", wspd.websocket.Clients)
	
	// There's no authentication for this, so you can end up with infinite connections.
	// Hadis => How would I do the auth? How can I comapre 2 ws conns?

	// if wspd.websocket.Clients[id][conn] {
	// 	return
	// }
	// Can I do something like this?
	// It seems that golang does a deep comparison instead of a shallow. So it would work?

	wspd.mutex.Unlock()

	go wspd.startWsProductDomain()

	for {
		var message service.ProductWsMessage

		if err := websocket.JSON.Receive(conn, &message); err != nil {
			conn.Close()
			wspd.removeWsProductClient(string(message.ID), conn)
	
			wspd.websocket.ErrorChan <- true

			return
		}

		wspd.websocket.BroadcastChan <- message
	}
	
}

func (wspd *WsProductDomain) removeWsProductClient(id string, conn *websocket.Conn) {
	wspd.mutex.Lock()

	delete(wspd.websocket.Clients[id], conn)

	if (len(wspd.websocket.Clients[id]) == 0) {
		delete(wspd.websocket.Clients, id)
	}

	wspd.mutex.Unlock()
}

func (wspd *WsProductDomain) startWsProductDomain() {
	for {
		select {
		case message := <-wspd.websocket.BroadcastChan:
			wspd.HandleWsProductBroadcastMsg(message)
		case err := <- wspd.websocket.ErrorChan:
			if err {
				return
			}
		}
	}
}

func (wspd *WsProductDomain) HandleWsProductBroadcastMsg(msg service.ProductWsMessage) {
	fmt.Println("sasasas", msg.ID, msg.ID)
	//fmt.Println("sasasas2", wspd.websocket.Clients)

	clients := wspd.websocket.Clients[msg.ID]
	
	fmt.Println("wspd.websocket.Clients", wspd.websocket.Clients)
	fmt.Println("clients", clients)

	for client := range clients {
		websocket.JSON.Send(client, msg)
	}
}
