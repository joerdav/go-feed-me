package types

type Restaurant struct {
	Id    string `schema:"id"`
	Name  string `schema:"name"`
	Items []Item `schema:"items" json:"menu"`
}

type Item struct {
	Id       int    `schema:"id"`
	Name     string `schema:"name" json:"item"`
	Price    int    `schema:"price"`
	Quantity int    `schema:"quantity"`
}

type Basket struct {
	Restaurants []Restaurant
}
