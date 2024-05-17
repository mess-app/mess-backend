package ws

import (
	"errors"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/nedpals/supabase-go"
)

type WebsocketHub struct {
	sync.RWMutex
	// websocket connection of every connected user. Key: auth.users.id
	connections map[string]*websocket.Conn
	clients     map[string]*supabase.Client
}

func NewWebsocketHub() *WebsocketHub {
	return &WebsocketHub{}
}

func (hub *WebsocketHub) addConnection(userId string, conn *websocket.Conn) {
	hub.Lock()
	hub.connections[userId] = conn
	hub.Unlock()
}

func (hub *WebsocketHub) removeConnection(userId string) {
	hub.Lock()
	delete(hub.connections, userId)
	hub.Unlock()
}

func (hub *WebsocketHub) getConnection(userId string) (*websocket.Conn, error) {
	hub.RLock()
	conn, ok := hub.connections[userId]

	if !ok {
		return nil, errors.New("connection not found")
	}

	defer hub.RUnlock()
	return conn, nil
}

func (hub *WebsocketHub) addClient(userId string, client *supabase.Client) {
	hub.Lock()
	hub.clients[userId] = client
	hub.Unlock()
}

func (hub *WebsocketHub) removeClient(userId string) {
	hub.Lock()
	delete(hub.clients, userId)
	hub.Unlock()
}

func (hub *WebsocketHub) getClient(userId string) (*supabase.Client, error) {
	hub.RLock()
	client, ok := hub.clients[userId]

	if !ok {
		return nil, errors.New("client not found")
	}

	defer hub.RUnlock()
	return client, nil
}
