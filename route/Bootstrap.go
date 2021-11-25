package route

import "github.com/gorilla/mux"

func Bootstrap(r *mux.Router) {

	todos(r)
}
