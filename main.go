package main

import (
	"fmt"
	"log"
	"net/http"
)

// GET all Users

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Output(1, "Request Sent")
}

func main() {
	fmt.Println("Hi From Main ")

}
