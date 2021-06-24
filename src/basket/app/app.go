package app

import (
	"basket/types"
	"fmt"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var inmemorybasket = map[string]types.Restaurant{
	"1": types.Restaurant{
		Id:   "1",
		Name: "Becky's Burgers",
		Items: []types.Item{
			types.Item{
				Id:       0,
				Name:     "Cheeseburger",
				Price:    9,
				Quantity: 1,
			},
			types.Item{
				Id:       1,
				Name:     "Milkshake",
				Price:    4,
				Quantity: 0,
			},
			types.Item{
				Id:       2,
				Name:     "Meal (burger, fries, and shake)",
				Price:    14,
				Quantity: 0,
			},
		},
	},
}

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
		PathPrefix("/apps").
		Subrouter()

	r.Handle("/basket", UpdateBasketHandler{Config: c}).Methods("PUT")
	r.HandleFunc("/basket", HandleGetBasketRequest).Methods("GET")

	corsObj := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"PUT", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{"turbo-frame"})

	fmt.Printf("Listening: %s", c.Port)

	err := http.ListenAndServe(c.Port, handlers.CORS(corsObj, headers, methods)(r))

	log.Fatal(err)
}
