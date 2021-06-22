package app

import (
	"basket/templates"
	"basket/types"
	"context"
	"net/http"

	"github.com/joe-davidson1802/hotwirehandler"
	"github.com/joe-davidson1802/turbo-templ/turbo"
)

type GetBasketHandler struct {
	Config types.Config
}

func (h GetBasketHandler) CanHandleModel(m string) bool {
	return m == types.Basket{}.ModelName()
}

func (h GetBasketHandler) HandleRequest(w http.ResponseWriter, r *http.Request) (error, hotwirehandler.Model) {
	basket := []types.Restaurant{}

	for _, res := range inmemorybasket {
		basket = append(basket, res)
	}

	w.Header().Add("Cache-Control", "no-cache")

	return nil, types.Basket{
		Restaurants: basket,
	}
}

func (h GetBasketHandler) RenderPage(ctx context.Context, m hotwirehandler.Model, w http.ResponseWriter) error {
	mod := m.(types.Basket)

	w.Header().Add("Content-Type", "text/html")

	contents := templates.BasketComponent(mod)

	frame := turbo.TurboFrame(turbo.TurboFrameOptions{
		Id:       "basket",
		Contents: &contents,
	})

	err := frame.Render(ctx, w)

	return err
}

func (h GetBasketHandler) RenderStream(ctx context.Context, m hotwirehandler.Model, w http.ResponseWriter) error {
	mod := m.(types.Basket)

	w.Header().Add("Content-Type", "text/vnd.turbo-stream.html")

	contents := templates.BasketComponent(mod)

	stream := turbo.TurboStream(turbo.TurboStreamOptions{
		Action:   turbo.Update,
		Target:   "basket",
		Contents: &contents,
	})

	err := stream.Render(ctx, w)

	return err
}
