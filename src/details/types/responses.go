package types

type Restaurant struct {
	Id               string
	Name             string
	PriceRange       string
	ImageSrc         string
	ImageDescription string
	Description      string
	Menu             []Item
}

func (m Restaurant) ModelName() string { return "Restaurant" }

type Item struct {
	Name  string `json:"item"`
	Price int
}

type RestaurantList struct {
	Restaurants []Restaurant
}

func (m RestaurantList) ModelName() string { return "RestaurantList" }
