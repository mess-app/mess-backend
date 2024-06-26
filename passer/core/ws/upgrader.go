package ws

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/KRTirtho/mess-backend/passer/core/collections"
	"github.com/KRTirtho/mess-backend/passer/core/collections/tables"
	"github.com/KRTirtho/mess-backend/passer/core/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/nedpals/supabase-go"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (hub *WebsocketHub) HandleWebsocketConnection(context *gin.Context) {
	authorizationHeader, ok := context.GetQuery("token")

	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	client := supabase.CreateClient(
		collections.Env.SupabaseURL,
		collections.Env.SupabaseAnonKey,
	)

	user, err := client.Auth.User(context, authorizationHeader)

	client.DB.AddHeader("Authorization", "Bearer "+authorizationHeader)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return

	}

	profile := models.SupabaseProfileId{}

	err = client.DB.From(tables.Profile).
		Select("id").
		Single().
		Eq("user_id", user.ID).
		Execute(&profile)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	context.Set("profileId", profile.Id)
	context.Set("userId", user.ID)

	connection, err := upgrader.Upgrade(context.Writer, context.Request, nil)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	defer func() {
		hub.removeConnection(user.ID)
		hub.removeClient(user.ID)
		connection.Close()
	}()

	hub.addConnection(user.ID, connection)
	hub.addClient(user.ID, client)

	for {
		// Reading All incoming Websocket messages
		messageType, message, err := connection.ReadMessage()

		if err != nil {
			fmt.Println(err)
			break
		}

		if messageType == websocket.TextMessage {

			jsonMap := map[string]interface{}{}

			err = json.Unmarshal(message, &jsonMap)

			if err != nil {
				fmt.Println(err)
				continue
			}

			data, ok := jsonMap["data"].(map[string]interface{})

			if !ok {
				fmt.Println("Invalid event data from ", user.ID)
				continue
			}

			event := models.WsEvent{
				Event: jsonMap["event"].(string),
				Data:  data,
			}

			hub.OnWebsocketEvent(context, event)
		} else {
			fmt.Println("Unknown message type from ", user.ID)
		}

	}
}
