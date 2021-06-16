package app

import (
	"context"
	"details/restaurants"
	"details/templates"
	"details/types"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joe-davidson1802/hotwirehandler"
)

type RestaurantHandler struct {
	Config types.Config
}

func (h RestaurantHandler) CanHandleModel(m string) bool {
	return m == types.Restaurant{}.ModelName()
}

func (h RestaurantHandler) HandleRequest(w http.ResponseWriter, r *http.Request) (error, hotwirehandler.Model) {
	vars := mux.Vars(r)
	id := vars["id"]

	repo := restaurants.RestaurantRepository{Config: h.Config}

	rs, err := repo.GetRestaurants()

	if err != nil {
		return err, nil
	}

	var restaurant types.Restaurant

	for _, r := range rs {
		if r.Id == id {
			restaurant = r
		}
	}

	return nil, restaurant
}

func (h RestaurantHandler) RenderPage(ctx context.Context, m hotwirehandler.Model, w http.ResponseWriter) error {
	mod := m.(types.Restaurant)

	w.Header().Add("Content-Type", "text/html")

	err := templates.RestaurantComponent(h.Config, mod).Render(ctx, w)

	return err
}

func (h RestaurantHandler) RenderStream(ctx context.Context, m hotwirehandler.Model, w http.ResponseWriter) error {
	return errors.New("Endpoint does not render streams")
}
