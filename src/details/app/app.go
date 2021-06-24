package app

import (
	"details/types"
	"fmt"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Run(env string) {
	var c types.Config

	config := "local.toml"

	if env != "" {
		config = env + ".toml"
	}

	if _, err := toml.DecodeFile(config, &c); err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter().
		PathPrefix("/apps/details").
		Subrouter()

	r.Handle("/restaurant/{id}", RestaurantHandler{Config: c}).Methods("GET", "OPTIONS")

	corsObj := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{"turbo-frame"})

	fmt.Printf("Listening: %s", c.Port)

	err := http.ListenAndServe(c.Port, handlers.CORS(corsObj, headers, methods)(r))

	log.Fatal(err)
}
