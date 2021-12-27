package route

import (
	"mux-training/auth"

	"github.com/gorilla/mux"
)

func common(r *mux.Router) {
	common := r.PathPrefix("/app").Subrouter()

	// Login
	common.HandleFunc("/login", auth.OauthGoogleLogin).Methods("GET")

	// Call back
	common.HandleFunc("/callback", auth.OauthGoogleCallback).Methods("GET")

}
