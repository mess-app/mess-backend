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
	return &WebsocketHub{
		connections: make(map[string]*websocket.Conn),
		clients:     make(map[string]*supabase.Client),
	}
}

func (hub *WebsocketHub) addConnection(userId string, conn *websocket.Conn) {
	hub.Lock()
	defer hub.Unlock()
	hub.connections[userId] = conn
}

func (hub *WebsocketHub) removeConnection(userId string) {
	hub.Lock()
	defer hub.Unlock()
	delete(hub.connections, userId)
}

func (hub *WebsocketHub) getConnection(userId string) (*websocket.Conn, error) {
	hub.RLock()
	defer hub.RUnlock()

	conn, ok := hub.connections[userId]

	if !ok {
		return nil, errors.New("connection not found")
	}

	return conn, nil
}

func (hub *WebsocketHub) addClient(userId string, client *supabase.Client) {
	hub.Lock()
	defer hub.Unlock()

	hub.clients[userId] = client
}

func (hub *WebsocketHub) removeClient(userId string) {
	hub.Lock()
	delete(hub.clients, userId)
	hub.Unlock()
}

func (hub *WebsocketHub) getClient(userId string) (*supabase.Client, error) {
	hub.RLock()
	defer hub.RUnlock()

	client, ok := hub.clients[userId]

	if !ok {
		return nil, errors.New("client not found")
	}

	return client, nil
}
