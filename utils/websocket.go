package utils

import (
	"encoding/json"
	"fmt"
	"real-time-forum/variables"
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
	nickname, ok := h.clients[conn]
	if ok {
		// Diffuse le message de déconnexion AVANT de supprimer
		message := fmt.Sprintf("%s s'est déconnecté", nickname)
		h.mu.Unlock() // On unlock ici pour éviter le deadlock avec BroadcastMessage
		h.BroadcastMessage([]byte(message))

		h.mu.Lock() // On relock pour continuer à modifier les clients
		delete(h.clients, conn)
	}
	h.mu.Unlock()
	conn.Close()
}


// BroadcastMessage envoie un message à tous les clients connectés
func (h *Hub) BroadcastMessage(message []byte) {
    h.mu.Lock()
    defer h.mu.Unlock()
    
    // Conversion en string puis création d'un JSON
    messageString := string(message)
    jsonMessage, _ := json.Marshal(map[string]string{"type": "log", "content": messageString})
    
    for conn := range h.clients {
        err := conn.WriteMessage(websocket.TextMessage, jsonMessage)
        if err != nil {
            conn.Close()
            delete(h.clients, conn)
        }
    }
}

func (h *Hub) SendMessage(message variables.Message){
    h.mu.Lock()
    defer h.mu.Unlock()
    
    // Conversion en string puis création d'un JSON
    // messageString := string(message)
    // jsonMessage, _ := json.Marshal(map[string]string{"type": "message", "content": messageString, "sender": sender, "receiver": receiver})
    
    for conn := range h.clients {
		if (h.clients[conn] != message.Receiver) {
			continue
		}
        err := conn.WriteJSON(message)
        if err != nil {
            conn.Close()
            delete(h.clients, conn)
        }
		fmt.Println("Message sent to", h.clients[conn], ":", message.Content)
    }
}

// func (h *Hub) GetOnlineUsers() []string {
//     h.mu.Lock()
//     defer h.mu.Unlock()
    
//     onlineUsers := make([]string, 0, len(h.clients))
//     for _, nickname := range h.clients {
//         onlineUsers = append(onlineUsers, nickname)
//     }
    
//     fmt.Println("Online users:", onlineUsers)
//     return onlineUsers
// }

// func (h *Hub) GetHub() *Hub {
// 	return h
// }