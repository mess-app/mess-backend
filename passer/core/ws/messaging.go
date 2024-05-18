package ws

import (
	"github.com/KRTirtho/mess-backend/passer/core/models"
)

func (hub *WebsocketHub) OnWebsocketEvent(event models.WsEvent, userId string) {
	switch event.Event {
	case "send_message":
		hub.OnSendMessage(event, userId)
	default:
		break
	}
}
