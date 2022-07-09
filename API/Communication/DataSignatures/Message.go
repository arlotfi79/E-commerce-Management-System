package DataSignatures

import "time"

type Message struct {
	Id          uint64    `json:"id"`
	MessageText string    `json:"message_text"`
	MessageDate time.Time `json:"message_date"`
}
