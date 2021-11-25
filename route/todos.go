package route

import (
	"mux-training/handler"
	"mux-training/util"
	"mux-training/validator"

	"github.com/gorilla/mux"
)

func todos(r *mux.Router) {
	todoRouter := r.PathPrefix("/todos").Subrouter()

	// Get todo list
	todoRouter.HandleFunc("/", handler.TodoList).Methods("GET")

	// Create a todo
	todoRouter.HandleFunc("/", util.Chain(handler.TodoCreate, validator.TodoPayload())).Methods("POST")

	// Delete all todo
	todoRouter.HandleFunc("/", handler.TodoDeleteAll).Methods("DELETE")

	// Delete todo by ID
	todoRouter.HandleFunc("/{id}", util.Chain(handler.TodoDeleteByID, validator.TodoCheckExistance())).Methods("DELETE")

	// Update todo by id
	todoRouter.HandleFunc("/{id}", util.Chain(handler.TodoUpdateByID, validator.TodoPayload(), validator.TodoCheckExistance())).Methods("PATCH")

}
