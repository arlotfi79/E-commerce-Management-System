package DataSignatures

type GetProduct struct {
	Id       uint64  `json:"id"`
	Name     string  `json:"name"`
	Color    string  `json:"color"`
	Price    float64 `json:"price"`
	Weight   float64 `json:"weight"`
	Quantity uint64  `json:"quantity"`
}

type GetProductFromCart struct {
	GetProduct
	RequestedQuantity uint64 `json:"requested_quantity"`
}
