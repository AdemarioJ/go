package main

import (
	"crud/server"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", server.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", server.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", server.ShowUser).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":5000", router))
}
