package app

import (
	"fmt"
	"log"
	"net/http"
	"random/types"

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

	r := mux.NewRouter()

	r.Handle("/apps/random", RandomHandler{Config: c}).Methods("GET", "OPTIONS")

	corsObj := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{"turbo-frame"})

	fmt.Printf("Listening: %s", c.Port)

	err := http.ListenAndServe(c.Port, handlers.CORS(corsObj, headers, methods)(r))

	log.Fatal(err)
}
