package main

import (
	"encoding/json"
	"log"
	"net/http"
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

}
