package app

import (
	"fmt"
	"log"
	"net/http"

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

func Run() {
	r := mux.NewRouter().
		PathPrefix("/content").
		Subrouter()

	r.
		PathPrefix("/").
		Handler(http.StripPrefix("/content/", http.FileServer(http.Dir("./public/"))))

	r.Use(loggingMiddleware)

	err := http.ListenAndServe(":8082", r)

	log.Fatal(err)
}
