package app

import (
	"details/restaurants"
	"details/templates"
	"details/types"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joe-davidson1802/turbo-templ/turbo"
)

type RestaurantHandler struct {
	Config types.Config
}

func (h RestaurantHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	repo := restaurants.RestaurantRepository{Config: h.Config}

	rs, err := repo.GetRestaurants()

	if err != nil {
		log.Fatal(err)
		return
	}

	var restaurant types.Restaurant

	for _, r := range rs {
		if r.Id == id {
			restaurant = r
		}
	}

	w.Header().Add("Cache-Control", "no-cache")
	w.Header().Add("Content-Type", "text/html")

	content := templates.RestaurantComponent(h.Config, restaurant)

	frame := turbo.TurboFrame(turbo.TurboFrameOptions{
		Id:       "container",
		Contents: &content,
	})

	err = frame.Render(r.Context(), w)

	if err != nil {
		log.Fatal(err)
	}
}
