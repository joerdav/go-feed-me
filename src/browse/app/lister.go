package app

import (
	"browse/restaurants"
	"browse/templates"
	"browse/types"
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/joe-davidson1802/hotwirehandler"
)

type ListerHandler struct {
	Config types.Config
}

func (h ListerHandler) CanHandleModel(m string) bool {
	return m == types.RestaurantList{}.ModelName()
}

func (h ListerHandler) HandleRequest(w http.ResponseWriter, r *http.Request) (error, hotwirehandler.Model) {
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
		return err, nil
	}

	return nil, types.RestaurantList{
		Restaurants: resultList,
	}
}

func (h ListerHandler) RenderPage(ctx context.Context, m hotwirehandler.Model, w http.ResponseWriter) error {
	mod := m.(types.RestaurantList)

	w.Header().Add("Content-Type", "text/html")

	err := templates.ListerComponent(h.Config, mod.Restaurants).Render(ctx, w)

	return err
}

func (h ListerHandler) RenderStream(ctx context.Context, m hotwirehandler.Model, w http.ResponseWriter) error {
	return errors.New("Endpoint does not render streams")
}
