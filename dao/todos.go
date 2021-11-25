package dao

import (
	"database/sql"
	"errors"
	"mux-training/model"
	"mux-training/module/sqlite"
)

func TodoList(title string, status interface{}) (todos []model.Todo, err error) {
	// Get database
	db := sqlite.GetDatabase()

	// Prepare query
	var stt bool
	q := "select * from todos where title like ?"
	if status != nil {
		stt = status.(bool)
		q += " and status = ?"
	}

	// Call query
	var rows *sql.Rows
	if status != nil {
		rows, err = db.Query(q, title, stt)
	} else {
		rows, err = db.Query(q, title)
	}

	if err != nil {
		return
	}

	// Bind data to todo list
	for rows.Next() {
		var todo model.Todo
		rows.Scan(&todo.ID, &todo.Title, &todo.Status)
		todos = append(todos, todo)
	}

	return
}

func TodoCreate(payload model.TodoPayload) (todoID int, err error) {
	// Get database
	db := sqlite.GetDatabase()

	// Execute statement
	rs, err := db.Exec("insert into todos(title, status) values (?, ?)", payload.Title, payload.Status)

	if err != nil {
		return
	}

	// Get new ID for todo
	newID, err := rs.LastInsertId()

	if err != nil {
		return
	}

	todoID = int(newID)

	return
}

func TodoDeleteByID(id int) error {
	// Get database
	db := sqlite.GetDatabase()

	// Execute statement
	_, err := db.Exec("delete from todos where id = ?", id)

	return err
}

func TodoDeleteAll() error {
	// Get database
	db := sqlite.GetDatabase()

	// Execute statement
	_, err := db.Exec("delete from todos")

	return err
}

func TodoUpdateByID(id int, payload model.TodoPayload) (err error) {
	// Get database
	db := sqlite.GetDatabase()

	// Execute statement
	_, err = db.Exec("update todos set title = ?, status = ? where id = ?", payload.Title, payload.Status, id)

	if err != nil {
		return
	}

	return
}

func TodoFindByID(id int) (todo model.Todo, err error) {

	// Get database
	db := sqlite.GetDatabase()

	// Call query
	rows, err := db.Query("select * from todos where id = ?", id)

	if err != nil {
		return
	}

	// Bind data to todo list
	for rows.Next() {
		rows.Scan(&todo.ID, &todo.Title, &todo.Status)
	}

	if todo.Title == "" {
		return todo, errors.New("Todo not found")
	}

	return
}
