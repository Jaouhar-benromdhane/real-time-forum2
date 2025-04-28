package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"real-time-forum/database"
	"real-time-forum/variables"
)

func Post(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Bienvenue sur la page de création de post !")

	user := database.GetCurrentUser(r)
	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	post := variables.Post{}
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	post.User = user
	database.InsertPost(&post)

}
