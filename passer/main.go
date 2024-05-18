package main

import (
	"log"

	"github.com/KRTirtho/mess-backend/passer/core/collections"
	"github.com/KRTirtho/mess-backend/passer/core/ws"
	"github.com/gin-gonic/gin"
	"github.com/nedpals/supabase-go"
)

// Supabase client with service role key
var SupabaseServerClient *supabase.Client

func main() {
	collections.Env = collections.NewEnv()

	client := supabase.CreateClient(collections.Env.SupabaseURL, collections.Env.SupabaseServiceKey, true)

	SupabaseServerClient = client

	router := gin.Default()

	hub := ws.NewWebsocketHub()

	go hub.Run()

	router.GET("/ws", hub.HandleWebsocketConnection)
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	log.Fatal(router.Run(collections.Env.Port))
}
