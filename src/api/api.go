package api

import (
	"encoding/json"
	"golang-rest-api/src/models"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Data struct {
	Success bool          `json:"success"`
	Data    []models.Todo `json:"data"`
	Errors  []string      `json:"errors"`
}

func CreateTodo(w http.ResponseWriter, req *http.Request) {
	var bodyTodo models.Todo
	err := json.NewDecoder(req.Body).Decode(&bodyTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var data Data = Data{Errors: make([]string, 0)}
	bodyTodo.Description = strings.TrimSpace(bodyTodo.Description)
	if len(bodyTodo.Description) == 0 {
		data.Success = false
		data.Errors = append(data.Errors, "invalid description")

		json, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(json)
		return
	}

	todo := models.Insert(bodyTodo.Description)

	data.Data = append(data.Data, todo)
	data.Success = true

	json, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(json)
	return
}

func GetTodos(w http.ResponseWriter, req *http.Request) {
	var todos []models.Todo = models.GetAll()

	var data = Data{true, todos, make([]string, 0)}

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
