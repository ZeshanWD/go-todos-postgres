package api

import (
	"golang-rest-api/src/database"
	"golang-rest-api/src/models"
	"io"
	"log"
	"net/http"
)

func CreateTodo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello world from server")
}

func GetTodos(res http.ResponseWriter, req *http.Request) {
	db := database.GetConnection()
	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {
		t := models.Todo{}
		err := rows.Scan(&t.ID, &t.description)
		if err != nil {
			log.Fatal(err)
		}

		t.description = strNull.String

		append(t, todos)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return todos
}

func UpdateTodo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "updating todos..")
}

func GetTodo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "get todo")
}

func DeleteTodo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "delete todo")
}
