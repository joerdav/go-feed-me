package app

import (
	"content/types"
	"fmt"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		fmt.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
		log.Println(r.RequestURI)
	})
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
		PathPrefix("/content").
		Subrouter()

	r.
		PathPrefix("/").
		Handler(http.StripPrefix("/content/", http.FileServer(http.Dir("./public/"))))

	r.Use(loggingMiddleware)

	err := http.ListenAndServe(c.Listen, r)

	log.Fatal(err)
}
