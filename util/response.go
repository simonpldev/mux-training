package util

import (
	"encoding/json"
	"net/http"
)

// Response
type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Response 200
func Response200(w http.ResponseWriter, message string, data interface{}) {
	if message == "" {
		message = "Success"
	}

	// Send OK
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Message: message,
		Data:    data,
	})
}

// Response 201
func Response201(w http.ResponseWriter, message string, data interface{}) {
	if message == "" {
		message = "Created"
	}

	// Send created
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Response{
		Message: message,
		Data:    data,
	})
}

// Response 204
func Response204(w http.ResponseWriter, message string) {
	if message == "" {
		message = "No content"
	}

	// Send no content
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(Response{
		Message: message,
	})
}

// Response 400
func Response400(w http.ResponseWriter, message string) {
	if message == "" {
		message = "Bad request"
	}

	// Send bad request
	http.Error(w, message, http.StatusBadRequest)
}

// Response 404
func Response404(w http.ResponseWriter, message string) {
	if message == "" {
		message = "Not found"
	}

	// Send not found
	http.Error(w, message, http.StatusNotFound)
}
