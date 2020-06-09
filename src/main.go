package main

import (
	"golang-rest-api/src/api"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/todos", api.GetTodos).Methods("GET")

	http.ListenAndServe(":8080", router)
}
