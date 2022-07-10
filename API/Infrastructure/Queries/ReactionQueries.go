package Queries

import (
	"API/Communication/DataSignatures"
	"API/Database"
	"log"
)

type ReactionQuery struct {
	dbClient *Database.Postgresql
}

func NewReactionQuery(dbClient *Database.Postgresql) *ReactionQuery {
	return &ReactionQuery{dbClient: dbClient}
}

// -------------------------------- POST ----------------------------------

func (reactionQuery *ReactionQuery) UpVoteReactionUsingAccountID(reaction *DataSignatures.VoteReaction) error {
	db := reactionQuery.dbClient.GetDB()

	query, err := db.Prepare(`INSERT INTO reaction (review_id, account_id, up_vote) 
									VALUES ($1, $2, $3)
									ON CONFLICT DO NOTHING`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	_, err = query.Exec(reaction.ReviewID, reaction.AccountID, reaction.Vote)

	if err != nil {
		return err
	}

	return nil
}

func (reactionQuery *ReactionQuery) DownVoteReactionUsingAccountID(reaction *DataSignatures.VoteReaction) error {
	db := reactionQuery.dbClient.GetDB()

	query, err := db.Prepare(`INSERT INTO reaction (review_id, account_id, down_vote) 
									VALUES ($1, $2, $3)
									ON CONFLICT DO NOTHING`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	_, err = query.Exec(reaction.ReviewID, reaction.AccountID, reaction.Vote)

	if err != nil {
		return err
	}

	return nil
}
