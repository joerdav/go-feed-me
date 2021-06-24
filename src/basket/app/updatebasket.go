package app

import (
	"basket/restaurants"
	"basket/templates"
	"basket/types"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/a-h/templ"
	"github.com/gorilla/schema"
	"github.com/joe-davidson1802/turbo-templ/turbo"
)

var decoder = schema.NewDecoder()

type UpdateBasketHandler struct {
	Config types.Config
}

func (h UpdateBasketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
		return
	}

	var restaurant types.Restaurant

	err = decoder.Decode(&restaurant, r.PostForm)

	if err != nil {
		log.Fatal(err)
		return
	}

	res := restaurants.RestaurantRepository{Config: h.Config}

	result, err := res.GetRestaurants()

	if err != nil {
		log.Fatal(err)
		return
	}

	id, err := strconv.Atoi(restaurant.Id)

	if err != nil {
		log.Fatal(err)
		return
	}

	restaurantdata := result[id-1]

	restaurant.Name = restaurantdata.Name

	for itemid, _ := range restaurant.Items {
		fmt.Println("ITEM")
		restaurant.Items[itemid].Price = restaurantdata.Items[itemid].Price
		restaurant.Items[itemid].Name = restaurantdata.Items[itemid].Name
		fmt.Println(restaurant.Items[itemid].Name)
	}

	basket := []types.Restaurant{}

	for _, res := range inmemorybasket {
		basket = append(basket, res)
	}

	basket = append(basket, restaurant)

	contents := templates.BasketComponent(types.Basket{
		Restaurants: basket,
	})

	w.Header().Add("Cache-Control", "no-cache")

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

	err = renderer.Render(r.Context(), w)

	if err != nil {
		log.Fatal(err)
	}
}
