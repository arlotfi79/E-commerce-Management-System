package DataSignatures

import "time"

type GetOrder struct {
	Id             uint64    `json:"id"`
	Description    string    `json:"description"`
	DeliveryMethod string    `json:"deliveryMethod"`
	OrderDate      time.Time `json:"orderDate"`
	Address        string    `json:"address"`
}

type PostOrder struct {
	Description    string    `json:"description"`
	DeliveryMethod string    `json:"deliveryMethod"`
	OrderDate      time.Time `json:"orderDate"`
	Address        string    `json:"address"`
}

type GetOrderTickets struct {
	Id             uint64    `json:"id"`
}