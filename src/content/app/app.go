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

	fmt.Println("reading config...")

	if _, err := toml.DecodeFile(config, &c); err != nil {
		log.Fatal(err)
	}

	fmt.Println("config loaded.")

	r := mux.NewRouter()

	r.
		PathPrefix("/content/").
		Handler(http.StripPrefix("/content/", http.FileServer(http.Dir("./public/"))))

	err := http.ListenAndServe(c.Listen, r)

	log.Fatal(err)
}
