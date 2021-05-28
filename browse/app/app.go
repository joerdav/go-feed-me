package app

import (
	"browse/hothandler"
	"browse/types"
	"fmt"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Run() {
	var c types.Config

	fmt.Println("reading config...")

	if _, err := toml.DecodeFile("dev.toml", &c); err != nil {
		log.Fatal(err)
	}

	fmt.Println("config loaded.")

	r := mux.NewRouter().
		PathPrefix("/browse")

	r.Handle("/restaurant/{id}", hothandler.New(RestaurantHandler{Config: c})).Methods("GET", "OPTIONS")
	r.Handle("/restaurants", hothandler.New(ListerHandler{Config: c})).Methods("GET", "OPTIONS")

	corsObj := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{"turbo-frame"})

	err := http.ListenAndServe(":8085", handlers.CORS(corsObj, headers, methods)(r))

	log.Fatal(err)
}
