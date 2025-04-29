package main

import (
	"fmt"
	"net/http"
	"real-time-forum/database"
	"real-time-forum/handler"
	"real-time-forum/utils"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {
	database.InitDb()
	hub := utils.NewHub()

	r := mux.NewRouter().StrictSlash(true)
	r.PathPrefix("/static/").Handler((http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))))
	r.HandleFunc("/", index).Methods("GET")
	//r.HandleFunc("/home", handler.Home).Methods("GET")
	r.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		handler.Home(w, r)
	})
	r.HandleFunc("/register", handler.Register).Methods("POST")
	r.HandleFunc("/login", handler.Login).Methods("POST")
	r.HandleFunc("/logout", handler.Logout).Methods("POST")
	r.HandleFunc("/post", handler.Post).Methods("POST")

	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handler.WebSocketHandler(w, r, hub)
	})

	r.HandleFunc("/refreshUsers", handler.RefreshUser).Methods("GET")

	r.HandleFunc("/comment", handler.CreateComment).Methods("POST")
	r.HandleFunc("/comment/{id}", handler.GetComments).Methods("GET")

	fmt.Printf("Server Started on http://localhost%s/\n", ":8080")
	http.ListenAndServe(":8080", r)
}

func index(w http.ResponseWriter, r *http.Request) {
	index := template.Must(template.ParseFiles("static/index.html"))
	index.Execute(w, nil)
}
