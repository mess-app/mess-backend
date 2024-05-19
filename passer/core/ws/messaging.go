package ws

import (
	"github.com/KRTirtho/mess-backend/passer/core/models"
	"github.com/gin-gonic/gin"
)

func (hub *WebsocketHub) OnWebsocketEvent(ctx *gin.Context, event models.WsEvent) {
	switch event.Event {
	case "send_message":
		hub.OnSendMessage(ctx, event)
	default:
		break
	}
}
