package handler

import (
	"net/http"
	"real-time-forum/database"
)

func Home(w http.ResponseWriter, r *http.Request) {
	user := database.GetCurrentUser(r)
	if user == nil {
		RespondJson(w, http.StatusUnauthorized, map[string]any{
			"error": "Unauthorized",
		})
		return
	}

	posts := database.GetpostHome()

	
		RespondJson(w, http.StatusOK, map[string]any{
			"Posts": posts,
		})
	
}

func RefreshUser(w http.ResponseWriter, r *http.Request){

	allUser := database.GetAllUsers(r)
	if len(allUser) == 0 {
		RespondJson(w, http.StatusNotFound, map[string]any{
			"error": "No Users Found",
		})
	} else {
		RespondJson(w, http.StatusOK, map[string]any{
			"Users": allUser,
		})
	}

}