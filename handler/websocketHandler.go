package handler

import (
	"fmt"
	"net/http"
	"real-time-forum/database"
	"real-time-forum/utils"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections; customize as needed
	},
}

type Message struct {
	Type     string `json:"type"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Content  string `json:"content"`
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request, hub *utils.Hub) {
	// Upgrade the connection to a WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()

	cookie, err := r.Cookie("session")

	if err != nil {
		fmt.Println("Error getting cookie:", err)
		return
	}
	session := database.GetUserIdBySession(cookie.Value)
	nickname := database.GetNicknameByUserId(session)

	hub.RegisterClient(conn, nickname)

	// Broadcast the message to all clients
	hub.BroadcastMessage([]byte(fmt.Sprintf("%s has joined the chat", nickname)))
	fmt.Println("Client connected:", nickname)

	// Listen for messages from the client
	for {
		var message Message
		err = (conn.ReadJSON(&message))
		if err != nil {
			fmt.Println("Error reading JSON:", err)
			return
		} else {
			if message.Type == "logout" || message.Type == "login" {
				hub.BroadcastMessage([]byte(message.Content))
			} else {
				hub.SendMessage([]byte(message.Content), message.Receiver, message.Sender)
			}
		}

	}

}
