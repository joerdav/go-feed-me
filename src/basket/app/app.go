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
