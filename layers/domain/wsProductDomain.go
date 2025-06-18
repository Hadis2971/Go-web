package domain

import (
	"sync"

	"golang.org/x/net/websocket"

	"github.com/Hadis2971/go_web/layers/service"
)

type WsProductDomain struct {
	websocket *service.WebsocketService[service.ProductMessage]
	mutex     sync.Mutex
}

func NewWsProductDomain(websocket *service.WebsocketService[service.ProductMessage]) *WsProductDomain {
	return &WsProductDomain{websocket: websocket}
}

func (wspd *WsProductDomain) AddNewWsProductClient(id string, conn *websocket.Conn) {

	wspd.mutex.Lock()

	if (wspd.websocket.Clients[id] == nil) {
		wspd.websocket.Clients[id] = make(map[*websocket.Conn]bool)
	}

	if (wspd.websocket.Clients[id][conn] == true) {
		return
	} else {
		wspd.websocket.Clients[id][conn] = true
	}
	
	// There's no authentication for this, so you can end up with infinite connections.
	// Hadis => How would I do the auth? How can I comapre 2 ws conns?

	// if wspd.websocket.Clients[id][conn] {
	// 	return
	// }
	// Can I do something like this?
	// It seems that golang does a deep comparison instead of a shallow. So it would work?

	wspd.mutex.Unlock()

	go wspd.startBroadcasePorductWsMsgs()
	go wspd.startReceiveProductWsMsgs(conn)	
	
	
}

func (wspd *WsProductDomain) removeWsProductClient(id string, conn *websocket.Conn) {
	wspd.mutex.Lock()

	delete(wspd.websocket.Clients[id], conn)

	if (len(wspd.websocket.Clients[id]) == 0) {
		delete(wspd.websocket.Clients, id)
	}

	wspd.mutex.Unlock()
}

func (wspd *WsProductDomain) startBroadcasePorductWsMsgs() {
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

func (wspd *WsProductDomain) startReceiveProductWsMsgs(conn *websocket.Conn) {
	for {
		var message service.ProductMessage

		if err := websocket.JSON.Receive(conn, &message); err != nil {
			conn.Close()
			wspd.removeWsProductClient(string(message.ID), conn)
	
			wspd.websocket.ErrorChan <- true

			return
		}

		wspd.websocket.BroadcastChan <- message
	}
}

func (wspd *WsProductDomain) HandleWsProductBroadcastMsg(msg service.ProductMessage) {
	clients := wspd.websocket.Clients[msg.ID]

	for client := range clients {
		websocket.JSON.Send(client, msg)
	}
}
