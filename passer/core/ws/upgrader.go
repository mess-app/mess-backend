package ws

import (
	"net/http"

	"github.com/KRTirtho/mess-backend/passer/core/collections"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/nedpals/supabase-go"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (hub *WebsocketHub) HandleWebsocketConnection(
	context *gin.Context,
) {
	authorizationHeader := context.GetHeader("Authorization")

	if authorizationHeader == "" {
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

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return

	}

	connection, err := upgrader.Upgrade(context.Writer, context.Request, nil)

	if err != nil {
		return
	}

	defer func() {
		hub.removeConnection(user.ID)
		hub.removeClient(user.ID)
		connection.Close()
	}()

	hub.addConnection(user.ID, connection)
	hub.addClient(user.ID, client)
}
