package app

import (
	"content/types"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/mux"
)

func logHandler(han http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		x, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}
		log.Println(fmt.Sprintf("%q", x))
		rec := httptest.NewRecorder()
		han.ServeHTTP(rec, r)
		log.Println(fmt.Sprintf("%q", rec.Body))
	}
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
		Handler(logHandler(http.StripPrefix("/content", http.FileServer(http.Dir("/src/public/")))))

	err := http.ListenAndServe(c.Listen, r)

	log.Fatal(err)
}
