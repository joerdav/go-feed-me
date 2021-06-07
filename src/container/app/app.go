package app

import (
	"container/hothandler"
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

	fmt.Println("reading config...")

	if _, err := toml.DecodeFile(config, &c); err != nil {
		log.Fatal(err)
	}

	fmt.Println("config loaded.")

	r := mux.NewRouter()

	r.PathPrefix("/").Handler(hothandler.New(HomeHandler{Config: c}))

	err := http.ListenAndServe(c.Listen, r)

	log.Fatal(err)
}
