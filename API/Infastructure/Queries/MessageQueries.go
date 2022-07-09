package Queries

import (
	"API/Communication/DataSignatures"
	"API/Database"
	"log"
	"time"
)

type MessageQuery struct {
	dbClient *Database.Postgresql
}

func NewMessageQuery(dbClient *Database.Postgresql) *MessageQuery {
	return &MessageQuery{dbClient: dbClient}
}

// -------------------------------- GET ----------------------------------

func (messageQuery *MessageQuery) GetMessagesByTicketID(id uint64) ([]DataSignatures.Message, error) {
	db := messageQuery.dbClient.GetDB()

	query, err := db.Prepare(`SELECT message_id, message_text, message_date
									FROM message
									WHERE ticket_id = $1`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	row, err := query.Query(id)

	if err != nil {
		log.Fatal(err)
	}

	var messages []DataSignatures.Message
	for row.Next() {
		var message DataSignatures.Message

		err = row.Scan(&message.Id, &message.MessageText, &message.MessageDate)

		if err != nil {
			log.Fatal(err)
		}
	}

	return messages, nil
}

// -------------------------------- POST ----------------------------------

func (messageQuery *MessageQuery) PostMessageUsingTicketID(ticketID uint64, messageText string) error {
	db := messageQuery.dbClient.GetDB()

	query, err := db.Prepare(`INSERT INTO message (ticket_id, message_text, message_date)
									VALUES ($1, $2, $3)`)
	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	_, err = query.Exec(ticketID, messageText, time.Now())

	if err != nil {
		return err
	}

	return nil
}
