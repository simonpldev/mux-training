package util

import (
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Log URI
		log.Println(r.RequestURI)

		next.ServeHTTP(w, r)
	})
}
