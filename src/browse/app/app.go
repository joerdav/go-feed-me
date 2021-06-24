package app

import (
	"browse/types"
	"fmt"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/motemen/go-loghttp/global"
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
		PathPrefix("/apps/browse").
		Subrouter()

	r.Handle("/restaurants", ListerHandler{Config: c}).Methods("GET", "OPTIONS")

	corsObj := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{"turbo-frame"})

	fmt.Printf("Listening: %s", c.Port)

	err := http.ListenAndServe(c.Port, handlers.CORS(corsObj, headers, methods)(r))

	log.Fatal(err)
}
