package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"real-time-forum/database"
	"real-time-forum/variables"
	"strconv"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {
	user := database.GetCurrentUser(r)
	if user == nil {
		http.Error(w, "Utilisateur non connecté", http.StatusUnauthorized)
		return
	}

	var comment variables.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}

	comment.UserID = user.ID
	database.InsertComment(&comment)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comment)
}

func GetComments(w http.ResponseWriter, r *http.Request) {
	postIDParam := r.URL.Path[len("/comment/"):]
	if postIDParam == "" {
		http.Error(w, "id manquant", http.StatusBadRequest)
		return
	}
	fmt.Println(postIDParam)
	postID, _ := strconv.Atoi(postIDParam)
	comments := database.GetCommentsByPostID(postID)
	fmt.Println(comments)
	if len(comments) == 0 {
		RespondJson(w, http.StatusNotFound, map[string]any{
			"error": "Aucun commentaire trouvé pour ce post",
		})
		return
	}
	RespondJson(w, http.StatusOK, map[string]any{
		"comments": comments,
	})
}
