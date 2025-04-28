package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
	"real-time-forum/database"
	"real-time-forum/variables"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type LoginResponse struct {
	Identifiant string `json:"identifiant"`
	Password    string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Bienvenue sur la page de connexion !")
	var user *variables.User
	userLogin := LoginResponse{}
	err := json.NewDecoder(r.Body).Decode(&userLogin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	email, err := mail.ParseAddress(userLogin.Identifiant)
	if err != nil {
		user = database.GetUserByNickname(userLogin.Identifiant)
	} else {
		user = database.GetUserByEmail(email.Address)
	}
	if user.Email == "" {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(userLogin.Password))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	} else {
		setCookie(w, user)
		fmt.Println("Login successful")
	}
}

func setCookie(w http.ResponseWriter, user *variables.User) {
    session_token, _ := uuid.NewV4()
	// expireAt := time.Now().Add(time.Hour * 1)
    cookie := http.Cookie{
        Name:     "session",
        Value:    session_token.String(),//base64.StdEncoding.EncodeToString(session_token.Bytes()),
        Path:     "/",
		MaxAge:  3600,
        HttpOnly: true,
    }
    http.SetCookie(w, &cookie)
	database.InsertSession(session_token.String(), user)
}


func GetCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Error(w, "No cookie found", http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "Cookie found:", cookie.Value)

	w.Write([]byte("cookie found!"))
}
