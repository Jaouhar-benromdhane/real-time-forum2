package utils

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type Hub struct {
	clients map[*websocket.Conn]string
	mu sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[*websocket.Conn]string),
	}
}
func (h *Hub) RegisterClient(conn *websocket.Conn, nickname string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.clients[conn] = nickname
	
}
func (h *Hub) UnregisterClient(conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.clients, conn)
    h.BroadcastMessage([]byte(fmt.Sprintf("%s has left the chat", h.clients[conn])))
}

// BroadcastMessage envoie un message à tous les clients connectés
func (h *Hub) BroadcastMessage(message []byte) {
    h.mu.Lock()
    defer h.mu.Unlock()
    
    // Conversion en string puis création d'un JSON
    messageString := string(message)
    jsonMessage, _ := json.Marshal(map[string]string{"connexion": messageString})
    
    for conn := range h.clients {
        err := conn.WriteMessage(websocket.TextMessage, jsonMessage)
        if err != nil {
            conn.Close()
            delete(h.clients, conn)
        }
    }
}


func (h *Hub) GetOnlineUsers() []string {
    h.mu.Lock()
    defer h.mu.Unlock()
    
    onlineUsers := make([]string, 0, len(h.clients))
    for _, nickname := range h.clients {
        onlineUsers = append(onlineUsers, nickname)
    }
    
    fmt.Println("Online users:", onlineUsers)
    return onlineUsers
}

// func (h *Hub) GetHub() *Hub {
// 	return h
// }