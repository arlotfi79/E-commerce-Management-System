package DataSignatures

import "time"

type PromotionCode struct {
	Id         uint64    `json:"id"`
	Value      float64   `json:"value"`
	ExpireDate time.Time `json:"expire_date"`
}
