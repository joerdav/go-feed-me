package app

import (
	"basket/types"
	"fmt"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joe-davidson1802/hotwirehandler"
)

var inmemorybasket = map[string]types.Restaurant{}

func Run(env string) {
	var c types.Config

	config := "local.toml"

	if env != "" {
		config = env + ".toml"
	}

	fmt.Println("reading config...")

	if _, err := toml.DecodeFile(config, &c); err != nil {
		log.Fatal(err)
	}

	fmt.Println("config loaded.")

	r := mux.NewRouter().
		PathPrefix("/apps").
		Subrouter()

	r.Handle("/basket", hotwirehandler.New(UpdateBasketHandler{Config: c})).Methods("PUT")
	r.Handle("/basket", hotwirehandler.New(GetBasketHandler{Config: c})).Methods("GET")

	corsObj := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"PUT", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{"turbo-frame"})

	err := http.ListenAndServe(c.Listen, handlers.CORS(corsObj, headers, methods)(r))

	log.Fatal(err)
}
