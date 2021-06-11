package restaurants

import (
	"browse/types"
	"encoding/json"
	"fmt"
	"net/http"
)

type RestaurantRepository struct {
	Config types.Config
}

func (r RestaurantRepository) GetRestaurants() ([]types.Restaurant, error) {
	var rs []types.Restaurant

	url := r.Config.ContentBaseUrl + "/restaurants.json"

	fmt.Printf("Getting restaurants from: %s", url)

	resp, err := http.Get(url)

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
