package Queries

import (
	"API/Communication/DataSignatures"
	"API/Database"
	"log"
)

type TicketQuery struct {
	dbClient *Database.Postgresql
}

func NewTicketQuery(dbClient *Database.Postgresql) *TicketQuery {
	return &TicketQuery{dbClient: dbClient}
}

// -------------------------------- GET ----------------------------------

func (ticketQuery *TicketQuery) GetTicketsByOrderID(id uint64) ([]DataSignatures.Ticket, error) {
	db := ticketQuery.dbClient.GetDB()

	query, err := db.Prepare(`SELECT ticket_id, subject, ticket_date
									FROM tickettracking
									WHERE order_id = $1`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	row, err := query.Query(id)

	if err != nil {
		log.Fatal(err)
	}

	var tickets []DataSignatures.Ticket
	for row.Next() {
		var ticket DataSignatures.Ticket

		err = row.Scan(&ticket.Id, &ticket.Subject, &ticket.TicketDate)

		if err != nil {
			log.Fatal(err)
		}
	}

	return tickets, nil
}

// -------------------------------- POST ----------------------------------

func (ticketQuery *TicketQuery) PostTicketUsingOrderID(orderID uint64, ticket *DataSignatures.Ticket) error {
	db := ticketQuery.dbClient.GetDB()

	query, err := db.Prepare(`INSERT INTO tickettracking (order_id, subject, ticket_date)
									VALUES ($1, $2, $3)`)
	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	_, err = query.Exec(orderID, ticket.Subject, ticket.TicketDate)

	if err != nil {
		return err
	}

	return nil
}
