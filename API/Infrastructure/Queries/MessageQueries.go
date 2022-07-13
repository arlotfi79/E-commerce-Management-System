package Queries

import (
	"API/Communication/DataSignatures"
	"API/Database"
	"log"
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

	// newer messages will be listed first
	query, err := db.Prepare(`SELECT message_id, message_text, message_date 
									FROM Message
									WHERE ticket_id = $1
									ORDER BY message_date DESC`)

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

func (messageQuery *MessageQuery) PostMessageUsingTicketID(message *DataSignatures.PostMessage) error {
	db := messageQuery.dbClient.GetDB()

	query, err := db.Prepare(`INSERT INTO Message (ticket_id, message_text, message_date)
									VALUES ($1, $2, $3)`)
	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	_, err = query.Exec(message.TicketID, message.MessageText, message.MessageDate)

	if err != nil {
		return err
	}

	return nil
}
