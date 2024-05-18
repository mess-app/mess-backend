package ws

import (
	"fmt"
	"runtime/debug"

	"github.com/KRTirtho/mess-backend/passer/core/collections/tables"
	"github.com/KRTirtho/mess-backend/passer/core/models"
	"github.com/KRTirtho/mess-backend/passer/core/utils"
	"github.com/mitchellh/mapstructure"
)

func (hub *WebsocketHub) OnSendMessage(event models.WsEvent, userId string) {
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

	connection := models.SupabaseConnection{}

	fmt.Println(map[string]interface{}{
		"id":         data.ConnectionId,
		"pioneer_id": userId,
		"status":     "connected",
	})

	profile := models.SupabaseProfileId{}

	err = client.DB.From(tables.Profile).
		Select("id").
		Single().
		Eq("user_id", userId).
		Execute(&profile)

	if err != nil {
		fmt.Println(err)
		debug.PrintStack()
		return
	}

	err = client.DB.
		From(tables.Connections).
		Select("*").
		Single().
		Eq("id", data.ConnectionId).
		Eq("pioneer_id", profile.Id).
		Eq("status", "connected").
		Execute(&connection)

	if err != nil {
		err = client.DB.
			From(tables.Connections).
			Select("*").
			Single().
			Eq("id", data.ConnectionId).
			Eq("recipient_id", profile.Id).
			Eq("status", "connected").
			Execute(&connection)

		if err != nil {
			fmt.Println(err)
			debug.PrintStack()
			return
		}
	}

	var recipientId string

	if connection.PioneerId == userId {
		recipientId = connection.RecipientId
	} else {
		recipientId = connection.PioneerId
	}

	recipientProfile := models.SupabaseProfileUserId{}

	err = client.DB.From(tables.Profile).
		Select("user_id").
		Single().
		Eq("id", recipientId).
		Execute(&recipientProfile)

	if err != nil {
		fmt.Println(err)
		debug.PrintStack()
		return
	}

	fmt.Println(recipientProfile)

	conn, err := hub.getConnection(recipientProfile.UserId)

	if err != nil {
		fmt.Println(err)
		debug.PrintStack()
		return
	}

	writeMessage, err := utils.EncodeStructToMap(data)

	if err != nil {
		fmt.Println(err)
		debug.PrintStack()
		return
	}

	conn.WriteJSON(models.WsEvent{
		Event: "receive_message",
		Data:  writeMessage,
	})

}
