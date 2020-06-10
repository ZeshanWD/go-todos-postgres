package api

import (
	"encoding/json"
	"golang-rest-api/src/models"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type Data struct {
	Success bool
	Data    []models.Todo
}

func CreateTodo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello world from server")
}

func GetTodos(w http.ResponseWriter, req *http.Request) {
	var todos []models.Todo = models.GetAll()

	var data = Data{true, todos}

	json, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func UpdateTodo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "updating todos..")
}

func GetTodo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var data Data

	var todo models.Todo
	var success bool
	todo, success = models.Get(id)
	if success != true {
		data.Success = false

		json, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(json)
		return
	}

	data.Success = true
	data.Data = append(data.Data, todo)

	json, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(json)
}

func DeleteTodo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "delete todo")
}
