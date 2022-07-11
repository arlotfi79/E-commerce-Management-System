package DataSignatures

import "time"

type Message struct {
	Id          uint64    `json:"id"`
	MessageText string    `json:"message_text"`
	MessageDate time.Time `json:"message_date"`
}

type PostMessage struct {
	TicketID    uint64    `json:"ticket_id"`
	MessageText string    `json:"message_text"`
	MessageDate time.Time `json:"message_date"`
}
