package app

import (
	"container/types"
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

	r.PathPrefix("/").Handler(ContainerHandler{Config: c})

	fmt.Printf("Listening: %s", c.Port)

	err := http.ListenAndServe(c.Port, r)

	log.Fatal(err)
}
