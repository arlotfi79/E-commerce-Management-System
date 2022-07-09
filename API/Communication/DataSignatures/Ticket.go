package DataSignatures

import "time"

type Ticket struct {
	Id         uint64    `json:"id"`
	Subject    string    `json:"subject"`
	TicketDate time.Time `json:"ticket_date"`
}
