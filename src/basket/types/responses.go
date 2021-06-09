package types

type Restaurant struct {
	Id    string `schema:"id"`
	Name  string `schema:"name"`
	Items []Item `schema:"items" json:"menu"`
}

func (m Restaurant) ModelName() string { return "Restaurant" }

type Item struct {
	Id       int    `schema:"id"`
	Name     string `schema:"name" json:"item"`
	Price    int    `schema:"price"`
	Quantity int    `schema:"quantity"`
}

type Basket struct {
	Restaurants []Restaurant
}

func (m Basket) ModelName() string { return "Basket" }
func (m Item) ModelName() string   { return "Item" }
