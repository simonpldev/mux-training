package handler

import (
	"mux-training/model"
	"mux-training/service"
	"mux-training/util"
	"net/http"

	"github.com/gorilla/context"
)

// Get todo list
func TodoList(w http.ResponseWriter, r *http.Request) {
	// Set content type
	w.Header().Set("Content-Type", "application/json")

	// Get search option
	queries := r.URL.Query()
	title := queries.Get("title")
	status := queries.Get("status")

	// Call get todo list service
	todos, err := service.TodoList(title, status)

	if err != nil {
		util.Response404(w, err.Error())
		return
	}

	// Deliver todo list
	util.Response200(w, "", todos)
}

// Create todo
func TodoCreate(w http.ResponseWriter, r *http.Request) {
	// Set content type
	w.Header().Set("Content-Type", "application/json")
	payload := context.Get(r, "payload").(model.TodoPayload)

	// Call create todo service
	todo, err := service.TodoCreate(payload)

	if err != nil {
		util.Response404(w, err.Error())
		return
	}

	// Send response with created todo
	util.Response201(w, "", todo)
}

// Delete todo by id
func TodoDeleteByID(w http.ResponseWriter, r *http.Request) {
	// Set content type
	w.Header().Set("Content-Type", "application/json")

	// Get id from context
	id := context.Get(r, "id").(int)

	// Call delete todo service
	err := service.TodoDeleteByID(id)

	if err != nil {
		util.Response404(w, err.Error())
		return
	}

	// Send response
	util.Response204(w, "")
}

// Delete todo list
func TodoDeleteAll(w http.ResponseWriter, r *http.Request) {
	// Set content type
	w.Header().Set("Content-Type", "application/json")

	// Call delete todo service
	err := service.TodoDeleteAll()

	if err != nil {
		util.Response404(w, err.Error())
		return
	}

	// Send response
	util.Response204(w, "")
}

// Update todo by id
func TodoUpdateByID(w http.ResponseWriter, r *http.Request) {
	// Set content type
	w.Header().Set("Content-Type", "application/json")

	// Get data from context
	id := context.Get(r, "id").(int)
	payload := context.Get(r, "payload").(model.TodoPayload)

	// Call update todo service
	todo, err := service.TodoUpdateByID(id, payload)

	if err != nil {
		util.Response404(w, err.Error())
		return
	}

	// Send response with updated todo
	util.Response200(w, "", todo)
}
