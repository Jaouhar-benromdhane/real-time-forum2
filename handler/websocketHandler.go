package handler

import (
	"fmt"
	"net/http"
	"real-time-forum/database"
	"real-time-forum/utils"
	"real-time-forum/variables"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections; customize as needed
	},
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
		var message variables.Message
		err = (conn.ReadJSON(&message))
		if err != nil {
			fmt.Println("Error reading JSON:", err)
			return
		} else {
			fmt.Println("Received message:", message)
			if message.Type == "logout" || message.Type == "login" {
				hub.BroadcastMessage([]byte(message.Content))
			} else {
				message.Sender = database.GetNicknameByUserId(session)
				database.InsertMessage(&message)
				hub.SendMessage(message)

			}
		}

	}

}
