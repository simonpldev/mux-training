package model

// Todo
type Todo struct {
	ID     int    `json:"_id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

type TodoPayload struct {
	Title  string `json:"title" valid:"required, stringlength(1|50)"`
	Status bool   `json:"status" valid:"optional"`
}
