package application

import (
	"github.com/Hadis2971/go_web/layers/domain"
	"golang.org/x/net/websocket"
)

type WebsocketRoutesHandler struct {
	chatDomain *domain.ChatDomain
}

func NewWebsocketRoutesHandler (chatDomain *domain.ChatDomain) *WebsocketRoutesHandler {
	return &WebsocketRoutesHandler{chatDomain: chatDomain}
}

func (wrh *WebsocketRoutesHandler) Handler (ws *websocket.Conn) {
	id := ws.Request().URL.Query().Get("id")

	wrh.chatDomain.AddNewClient(id, ws)

	
}