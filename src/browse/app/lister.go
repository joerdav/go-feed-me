package app

import (
	"browse/hothandler"
	"browse/restaurants"
	"browse/templates"
	"browse/types"
	"context"
	"errors"
	"net/http"
)

type ListerHandler struct {
	Config types.Config
}

func (h ListerHandler) CanHandleModel(m string) bool {
	return m == types.RestaurantList{}.ModelName()
}

func (h ListerHandler) HandleRequest(w http.ResponseWriter, r *http.Request) (error, hothandler.Model) {
	repo := restaurants.RestaurantRepository{Config: h.Config}

	rs, err := repo.GetRestaurants()

	if err != nil {
		return errors.New("Failed to get restaurants"), nil
	}

	return nil, types.RestaurantList{
		Restaurants: rs,
	}
}

func (h ListerHandler) RenderPage(ctx context.Context, m hothandler.Model, w http.ResponseWriter) error {
	mod := m.(types.RestaurantList)

	w.Header().Add("Content-Type", "text/html")

	err := templates.ListerComponent(h.Config, mod.Restaurants).Render(ctx, w)

	return err
}

func (h ListerHandler) RenderStream(ctx context.Context, m hothandler.Model, w http.ResponseWriter) error {
	return errors.New("Endpoint does not render streams")
}
