package app

import (
	"basket/templates"
	"basket/types"
	"log"
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"github.com/joe-davidson1802/turbo-templ/turbo"
)

func HandleGetBasketRequest(w http.ResponseWriter, r *http.Request) {
	basket := []types.Restaurant{}

	for _, res := range inmemorybasket {
		basket = append(basket, res)
	}

	w.Header().Add("Cache-Control", "no-cache")

	contents := templates.BasketComponent(types.Basket{
		Restaurants: basket,
	})

	var renderer templ.Component

	if strings.Contains(r.Header.Get("Accept"), "vnd.turbo-stream.html") {
		w.Header().Add("Content-Type", "text/vnd.turbo-stream.html")

		renderer = turbo.TurboStream(turbo.TurboStreamOptions{
			Action:   turbo.UpdateAction,
			Target:   "basket",
			Contents: &contents,
		})
	} else {
		w.Header().Add("Content-Type", "text/html")

		renderer = turbo.TurboFrame(turbo.TurboFrameOptions{
			Id:       "basket",
			Contents: &contents,
		})
	}

	err := renderer.Render(r.Context(), w)

	if err != nil {
		log.Fatal(err)
	}
}
