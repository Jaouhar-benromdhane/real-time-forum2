package handler

import (
	"fmt"
	"net/http"
	"real-time-forum/database"
)


func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session")

	database.DeleteSession(cookie.Value)
	//cookie.MaxAge = -1 // Supprime le cookie
	cookie.Value = ""
	http.SetCookie(w, cookie)
	fmt.Println("Cookie supprimé")
	fmt.Println("Vous êtes déconnecté !")
}

