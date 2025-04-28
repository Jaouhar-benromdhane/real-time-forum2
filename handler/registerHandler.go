package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"real-time-forum/database"
	"real-time-forum/variables"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RegisterResponse struct {
	Nickname  string `json:"nickname"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bienvenue sur la page d’inscription !")
	userInfo := RegisterResponse{}
	err := json.NewDecoder(r.Body).Decode(&userInfo)
	if err != nil {
		fmt.Println(err)
		return
	}

	uuid, err := uuid.NewV4()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	    // Définir le cookie

	cookie := http.Cookie{
        Name:     "session",
        Value:    uuid.String(),
        Path:     "/",
        MaxAge:   3600,
        HttpOnly: true,
        Secure:   true,
        SameSite: http.SameSiteLaxMode,
    }
    http.SetCookie(w, &cookie)

	
    w.Write([]byte("Inscription réussie et session créée !"))


	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(userInfo.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := &variables.User{
		ID:        uuid.String(),
		Nickname:  userInfo.Nickname,
		Age:       userInfo.Age,
		Gender:    userInfo.Gender,
		FirstName: userInfo.FirstName,
		LastName:  userInfo.LastName,
		Email:     userInfo.Email,
		Password:  bcryptPassword,
	}

	database.InsertUser(user)
}


