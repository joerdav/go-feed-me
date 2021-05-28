package app

import (
	"container/hothandler"
	"container/types"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/mux"
)

func Run() {
	r := mux.NewRouter()

	var c types.Config

	fmt.Println("reading config...")

	if _, err := toml.DecodeFile("dev.toml", &c); err != nil {
		log.Fatal(err)
	}

	fmt.Println("config loaded.")
	fmt.Println(strconv.Itoa(len(c.Apps)) + " apps found")

	for _, a := range c.Apps {
		fmt.Print(a.Name + " " + a.Url)
		if a.Default {
			fmt.Println(" Default")
		}
		fmt.Println()
	}

	r.Handle("/", hothandler.New(HomeHandler{Config: c}))

	err := http.ListenAndServe(":80", r)

	log.Fatal(err)
}
