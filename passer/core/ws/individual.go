package ws

import (
	"fmt"
	"runtime/debug"

	"github.com/KRTirtho/mess-backend/passer/core/collections/tables"
	"github.com/KRTirtho/mess-backend/passer/core/models"
	"github.com/KRTirtho/mess-backend/passer/core/utils"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

func (hub *WebsocketHub) OnSendMessage(ctx *gin.Context, event models.WsEvent) {
	userId := ctx.GetString("userId")
	profileId := ctx.GetString("profileId")

	var data models.WsEventSendMessageData

	err := mapstructure.Decode(event.Data, &data)

	if err != nil {
		fmt.Println(err)
		debug.PrintStack()
		return
	}

	client, err := hub.getClient(userId)

	if err != nil {
		fmt.Println(err)
		debug.PrintStack()
		return
	}

	supabaseConnection := models.SupabaseConnection{}

	err = client.DB.
		From(tables.Connections).
		Select("*, recipient:profile!public_connections_recipient_id_fkey(user_id)").
		Single().
		Eq("id", data.ConnectionId).
		Eq("pioneer_id", profileId).
		Eq("status", "connected").
		Execute(&supabaseConnection)

	if err != nil {
		err = client.DB.
			From(tables.Connections).
			Select("*, pioneer:profile!public_connections_pioneer_id_fkey(user_id)").
			Single().
			Eq("id", data.ConnectionId).
			Eq("recipient_id", profileId).
			Eq("status", "connected").
			Execute(&supabaseConnection)

		if err != nil {
			fmt.Println(err)
			debug.PrintStack()
			return
		}
	}

	var recipientId string

	if supabaseConnection.PioneerId == profileId {
		recipientId = supabaseConnection.Recipient.UserId
	} else {
		recipientId = supabaseConnection.Pioneer.UserId
	}

	recipientConnection, err := hub.getConnection(recipientId)

	if err != nil {
		fmt.Println(err)
		debug.PrintStack()
		return
	}

	msgMap, err := utils.EncodeStructToMap(data)

	if err != nil {
		fmt.Println(err)
		debug.PrintStack()
		return
	}

	recipientConnection.WriteJSON(models.WsEvent{
		Event: "receive_message",
		Data:  msgMap,
	})

}
