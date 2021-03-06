package Queries

import (
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

func (reactionQuery *ReactionQuery) UpVoteReaction(reviewID uint64, accountID uint64) error {
	db := reactionQuery.dbClient.GetDB()

	query, err := db.Prepare(`INSERT INTO reaction (review_id, account_id, up_vote)
									VALUES ($1, $2, TRUE)
									ON CONFLICT (review_id, account_id)
										DO
        									UPDATE SET up_vote = TRUE`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	if err != nil {
		log.Fatal(err)
	}

	_, err = query.Exec(reviewID, accountID)

	if err != nil {
		return err
	}

	return nil
}

func (reactionQuery *ReactionQuery) RemoveUpVoteOfAnAccount(reviewID uint64, accountID uint64) error {
	db := reactionQuery.dbClient.GetDB()

	query, err := db.Prepare(`UPDATE reaction
									SET up_vote = FALSE
									WHERE review_id = $1 AND account_id = $2`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	if err != nil {
		log.Fatal(err)
	}

	_, err = query.Exec(reviewID, accountID)

	if err != nil {
		return err
	}

	return nil
}

func (reactionQuery *ReactionQuery) DownVoteReaction(reviewID uint64, accountID uint64) error {
	db := reactionQuery.dbClient.GetDB()

	query, err := db.Prepare(`INSERT INTO reaction (review_id, account_id, down_vote)
									VALUES ($1, $2, TRUE)
									ON CONFLICT (review_id, account_id)
										DO
        									UPDATE SET down_vote = TRUE`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	_, err = query.Exec(reviewID, accountID)

	if err != nil {
		return err
	}

	return nil
}

func (reactionQuery *ReactionQuery) RemoveDownVoteOfAnAccount(reviewID uint64, accountID uint64) error {
	db := reactionQuery.dbClient.GetDB()

	query, err := db.Prepare(`UPDATE reaction
									SET down_vote = FALSE
									WHERE review_id = $1 AND account_id = $2`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	if err != nil {
		log.Fatal(err)
	}

	_, err = query.Exec(reviewID, accountID)

	if err != nil {
		return err
	}

	return nil
}
