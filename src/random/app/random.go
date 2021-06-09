package app

import (
	"fmt"
	"math/rand"
	"net/http"
	"random/restaurants"
	"random/types"
	"time"
)

type RandomHandler struct {
	Config types.Config
}

func (h RandomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	repo := restaurants.RestaurantRepository{Config: h.Config}

	rs, err := repo.GetRestaurants()

	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().Unix())

	url := fmt.Sprintf("/apps/details/restaurant/%s", rs[rand.Intn(len(rs))].Id)

	http.Redirect(w, r, url, http.StatusSeeOther)
}
