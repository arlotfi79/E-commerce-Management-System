package Queries

import (
	"API/Communication/DataSignatures"
	"API/Database"
	"log"
)

type ReplyQuery struct {
	dbClient *Database.Postgresql
}

func NewReplyQuery(dbClient *Database.Postgresql) *ReplyQuery {
	return &ReplyQuery{dbClient: dbClient}
}

// -------------------------------- GET ----------------------------------

func (replyQuery *ReplyQuery) GetRepliesByReviewID(id uint64) ([]DataSignatures.GetReply, error) {
	db := replyQuery.dbClient.GetDB()

	query, err := db.Prepare(`SELECT reply_id, reply_text
									FROM Review_Reply 
									WHERE review_id = $1`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	row, err := query.Query(id)

	if err != nil {
		log.Fatal(err)
	}

	var replies []DataSignatures.GetReply
	for row.Next() {
		var reply DataSignatures.GetReply

		err = row.Scan(&reply.Id, &reply.ReplyText)

		if err != nil {
			log.Fatal(err)
		}

		replies = append(replies, reply)
	}

	return replies, nil
}

// -------------------------------- POST ----------------------------------

func (replyQuery *ReplyQuery) PostReplyUsingReviewID(reply *DataSignatures.PostReply) error {
	db := replyQuery.dbClient.GetDB()

	query, err := db.Prepare(`INSERT INTO Review_Reply (review_id, reply_text)
									VALUES ($1, $2)`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	_, err = query.Exec(reply.ReviewID, reply.ReplyText)

	if err != nil {
		return err
	}

	return nil
}
