package service

import (
	"mux-training/dao"
	"mux-training/model"
	"strings"
)

func TodoList(title string, status string) ([]model.Todo, error) {
	// Resolve keyword for searching
	title = "%" + strings.Trim(title, " ") + "%"

	var stt interface{}

	switch strings.Trim(status, " ") {
	case "true":
		stt = true
	case "false":
		stt = false
	default:
		stt = nil
	}

	// Call db for todo list
	return dao.TodoList(title, stt)
}

func TodoCreate(payload model.TodoPayload) (todo model.Todo, err error) {
	// Resolve value for todo payload
	payload.Title = strings.Trim(payload.Title, " ")
	payload.Status = false

	// Call db to create todo
	todo.ID, err = dao.TodoCreate(payload)
	if err != nil {
		return
	}

	// Set data for todo
	todo.Title = payload.Title
	todo.Status = payload.Status

	return
}

func TodoDeleteByID(id int) error {
	// Call db to delete todo
	return dao.TodoDeleteByID(id)
}

func TodoDeleteAll() error {
	// Call db to delete todo list
	return dao.TodoDeleteAll()
}

func TodoUpdateByID(id int, payload model.TodoPayload) (todo model.Todo, err error) {

	// Call db to create todo
	err = dao.TodoUpdateByID(id, payload)
	if err != nil {
		return
	}

	// Set data for todo
	todo.ID = id
	todo.Title = payload.Title
	todo.Status = payload.Status

	return
}
