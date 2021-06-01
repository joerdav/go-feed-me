package restaurants

import (
	"details/types"
	"encoding/json"
	"net/http"
)

type RestaurantRepository struct {
	Config types.Config
}

func (r RestaurantRepository) GetRestaurants() ([]types.Restaurant, error) {
	var rs []types.Restaurant

	resp, err := http.Get(r.Config.ContentBaseUrl + "/restaurants.json")

	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&rs)

	if err != nil {
		return nil, err
	}

	return rs, err
}
