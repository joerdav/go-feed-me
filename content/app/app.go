package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Run() {
	r := mux.NewRouter()

	staticDir := "/public/"

	r.
		PathPrefix("/").
		Handler(http.FileServer(http.Dir("." + staticDir)))

	err := http.ListenAndServe(":80", r)

	log.Fatal(err)
}
