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
	log.SetFlags(0)

	collections.Env = collections.NewEnv()

	client := supabase.CreateClient(collections.Env.SupabaseURL, collections.Env.SupabaseServiceKey, true)

	SupabaseServerClient = client

	router := gin.Default()

	hub := ws.NewWebsocketHub()

	go hub.Run()

	router.GET("/ws", hub.HandleWebsocketConnection)

	log.Fatal(router.Run(collections.Env.Port))
}
