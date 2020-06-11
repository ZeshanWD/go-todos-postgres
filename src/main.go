package main

import (
	"fmt"
	"golang-rest-api/src/api"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var port string = "8080"

	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api/").Subrouter()
	apiRouter.HandleFunc("/todos", api.GetTodos).Methods("GET")
	apiRouter.HandleFunc("/todos/{id}", api.GetTodo).Methods("GET")
	apiRouter.HandleFunc("/todos", api.CreateTodo).Methods("POST")

	fmt.Printf("Server running at port %s", port)
	http.ListenAndServe(":"+port, router)
}
