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

type Item struct {
	Name  string `json:"item"`
	Price int
}

type RestaurantList struct {
	Restaurants []Restaurant
}
