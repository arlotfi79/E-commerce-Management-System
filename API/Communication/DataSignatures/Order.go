package DataSignatures

import "time"

type Order struct {
	Id             uint64    `json:"id"`
	Description    string    `json:"description"`
	DeliveryMethod string    `json:"deliveryMethod"`
	OrderDate      time.Time `json:"orderDate"`
	Address        string    `json:"address"`
}
