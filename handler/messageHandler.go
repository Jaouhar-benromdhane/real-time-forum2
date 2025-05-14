package handler

import (
	"encoding/json"
	"net/http"
	"real-time-forum/database"

	"github.com/gorilla/mux"
)

func GetMessageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nickname := vars["nickname"]
	user := database.GetCurrentUser(r)
	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Retrieve messages from the database for the given nickname
	messages, err := database.GetMessages(nickname,user.Nickname)
	if err != nil {
		http.Error(w, "Failed to retrieve messages", http.StatusInternalServerError)
		return
	}

	// Respond with the messages in JSON format
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(messages); err != nil {
		http.Error(w, "Failed to encode messages", http.StatusInternalServerError)
	}
}