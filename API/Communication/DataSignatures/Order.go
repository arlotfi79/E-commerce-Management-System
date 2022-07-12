package DataSignatures

import "time"

type GetOrder struct {
	Id             uint64    `json:"id"`
	Description    string    `json:"description"`
	DeliveryMethod string    `json:"deliveryMethod"`
	OrderDate      time.Time `json:"orderDate"`
	Address        string    `json:"address"`
	IsComplete     bool      `json:"is_complete"`
}

type PostOrder struct {
	AccountID      uint64    `json:"account_id"`
	Description    string    `json:"description"`
	DeliveryMethod string    `json:"deliveryMethod"`
	OrderDate      time.Time `json:"orderDate"`
	Address        string    `json:"address"`
}
