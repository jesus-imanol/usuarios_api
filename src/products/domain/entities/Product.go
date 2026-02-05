package entities

type Product struct {
	Id       int32   `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int32   `json:"quantity"`
}

func NewProduct(name string, price float64, quantity int32) *Product {
	return &Product{
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}
}
