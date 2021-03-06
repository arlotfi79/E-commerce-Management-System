package DataSignatures

type PostProduct struct {
	Name     string  `json:"name"`
	Color    string  `json:"color"`
	Price    float64 `json:"price"`
	Weight   float64 `json:"weight"`
	Quantity uint64  `json:"quantity"`
}

type PostProductCategory struct {
	ProductID  uint64 `json:"product_id"`
	CategoryID uint64 `json:"category_id"`
}

type PostProductID struct {
	Id uint64 `json:"id"`
}

type GetProduct struct {
	Id uint64 `json:"id"`
	PostProduct
}

type GetProductFromCart struct {
	GetProduct
	RequestedQuantity uint64 `json:"requested_quantity"`
	EachEntryCost     uint64 `json:"each_entry_cost"`
	TotalCartCost     uint64 `json:"total_cart_cost"`
}
