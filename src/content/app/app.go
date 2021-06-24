package app

import (
	"content/types"
	"fmt"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
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

	r.
		PathPrefix("/content/").
		Handler(http.StripPrefix("/content/", http.FileServer(http.Dir("./public/"))))

	fmt.Printf("Listening: %s", c.Port)

	err := http.ListenAndServe(c.Port, r)

	log.Fatal(err)
}
