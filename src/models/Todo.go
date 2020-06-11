package models

import (
	"golang-rest-api/src/database"
	"log"
)

type Todo struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

func Insert(description string) Todo {
	db := database.GetConnection()

	var todo_id int
	db.QueryRow("INSERT INTO todos(description) VALUES($1) RETURNING id", description).Scan(&todo_id)

	return Todo{todo_id, ""}
}

func Get(id string) (Todo, bool) {
	db := database.GetConnection()
	row := db.QueryRow("SELECT * FROM todos WHERE id = $1", id)

	var ID int
	var description string
	err := row.Scan(&ID, &description)
	if err != nil {
		return Todo{}, false
	}

	return Todo{ID, description}, true
}

func Update() {

}

func GetAll() []Todo {
	db := database.GetConnection()
	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		t := Todo{}

		var ID int
		var description string

		err := rows.Scan(&ID, &description)
		if err != nil {
			log.Fatal(err)
		}

		t.ID = ID
		t.Description = description

		todos = append(todos, t)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return todos
}
