package app

import (
	"browse/restaurants"
	"browse/templates"
	"browse/types"
	"log"
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"github.com/joe-davidson1802/turbo-templ/turbo"
)

type ListerHandler struct {
	Config types.Config
}

func (h ListerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	repo := restaurants.RestaurantRepository{Config: h.Config}

	search := r.FormValue("search")

	rs, err := repo.GetRestaurants()

	resultList := []types.Restaurant{}

	if search != "" {
		for _, res := range rs {
			if strings.Contains(
				strings.ToLower(res.Name),
				strings.ToLower(search),
			) {
				resultList = append(resultList, res)
			}
		}
	} else {
		resultList = rs
	}

	if err != nil {
		log.Fatal(err)
		return
	}

	w.Header().Add("Cache-Control", "no-cache")

	var renderer templ.Component

	content := templates.ListerComponent(h.Config, resultList)

	w.Header().Add("Content-Type", "text/html")

	renderer = turbo.TurboFrame(turbo.TurboFrameOptions{
		Id:       "container",
		Contents: &content,
	})

	err = renderer.Render(r.Context(), w)

	if err != nil {
		log.Fatal(err)
	}
}
