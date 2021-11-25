package main

import (
	"log"
	"net/http"

	"mux-training/module/sqlite"
	"mux-training/route"
	"mux-training/util"

	"github.com/gorilla/mux"
)

func init() {
	sqlite.Connect()
}

func main() {
	r := mux.NewRouter()

	r.Use(mux.CORSMethodMiddleware(r))

	r.Use(util.LoggingMiddleware)

	route.Bootstrap(r)

	log.Fatal(http.ListenAndServe(":8000", r))
}
