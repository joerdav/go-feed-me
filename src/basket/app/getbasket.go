package app

import (
	"basket/hothandler"
	"basket/templates"
	"basket/types"
	"context"
	"net/http"
)

type GetBasketHandler struct {
	Config types.Config
}

func (h GetBasketHandler) CanHandleModel(m string) bool {
	return m == types.Basket{}.ModelName()
}

func (h GetBasketHandler) HandleRequest(w http.ResponseWriter, r *http.Request) (error, hothandler.Model) {
	basket := []types.Restaurant{}

	for _, res := range inmemorybasket {
		basket = append(basket, res)
	}

	return nil, types.Basket{
		Restaurants: basket,
	}
}

func (h GetBasketHandler) RenderPage(ctx context.Context, m hothandler.Model, w http.ResponseWriter) error {
	mod := m.(types.Basket)

	w.Header().Add("Content-Type", "text/html")

	err := templates.BasketFrameComponent(mod).Render(ctx, w)

	return err
}

func (h GetBasketHandler) RenderStream(ctx context.Context, m hothandler.Model, w http.ResponseWriter) error {
	mod := m.(types.Basket)

	w.Header().Add("Content-Type", "text/vnd.turbo-stream.html")

	err := templates.BasketFrameComponent(mod).Render(ctx, w)

	return err
}
