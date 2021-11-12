package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var users []User

// User Struct ( Model)

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Group  *Group `json:"group"`
	Status bool   `json:"status"`
}

type Group struct {
}

// GET all Users

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Output(1, "Request Sent")
	json.NewEncoder(w).Encode(users)
}

func main() {

	// Init Router
	r := mux.NewRouter()

	users = append(users, User{ID: 1, Name: "Arjun", Email: "arjunbiyani@gmail.com", Group: &Group{}, Status: true})
	// Router Handlers / Endpoint

	r.HandleFunc("/api/user", getUser).Methods("GET")

	// http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
	//     fmt.Fprintf(w, "Welcome to my website!")
	// })

	// fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":80", r))
}
