package application

import (
	"net/http"

	"github.com/Hadis2971/go_web/layers/domain"
	"golang.org/x/net/websocket"
)

type WebsocketRoutesHandler struct {
	chatDomain *domain.ChatDomain
}

func NewWebsocketRoutesHandler (chatDomain *domain.ChatDomain) *WebsocketRoutesHandler {
	return &WebsocketRoutesHandler{chatDomain: chatDomain}
}

func (wrh *WebsocketRoutesHandler) Handler (conn *websocket.Conn) http.HandlerFunc {
	

	return func (w http.ResponseWriter, r *http.Request) {
		param := r.URL.Query().Get("id")

		wrh.chatDomain.AddNewClient(param, conn)
	}

	
}