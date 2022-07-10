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

func (reactionQuery *ReactionQuery) UpVoteReaction(reaction *DataSignatures.VoteReaction) error {
	db := reactionQuery.dbClient.GetDB()

	query, err := db.Prepare(`INSERT INTO reaction (review_id, account_id, up_vote)
									VALUES ($1, $2, 1)
									ON CONFLICT (review_id, account_id)
										DO
											UPDATE SET up_vote = reaction.up_vote + 1`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	if err != nil {
		log.Fatal(err)
	}

	_, err = query.Exec(reaction.ReviewID, reaction.AccountID)

	if err != nil {
		return err
	}

	return nil
}

func (reactionQuery *ReactionQuery) RemoveUpVoteOfAnAccount(reaction *DataSignatures.VoteReaction) error {
	db := reactionQuery.dbClient.GetDB()

	query, err := db.Prepare(`UPDATE reaction
									SET up_vote = reaction.up_vote - 1
									WHERE up_vote > 0 AND review_id = $1 AND account_id = $2`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	if err != nil {
		log.Fatal(err)
	}

	_, err = query.Exec(reaction.ReviewID, reaction.AccountID)

	if err != nil {
		return err
	}

	return nil
}

func (reactionQuery *ReactionQuery) DownVoteReaction(reaction *DataSignatures.VoteReaction) error {
	db := reactionQuery.dbClient.GetDB()

	query, err := db.Prepare(`INSERT INTO reaction (review_id, account_id, down_vote) 
									VALUES ($1, $2, 1)
									ON CONFLICT 
									    DO 
									        UPDATE SET down_vote = reaction.down_vote + 1`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	_, err = query.Exec(reaction.ReviewID, reaction.AccountID)

	if err != nil {
		return err
	}

	return nil
}

func (reactionQuery *ReactionQuery) RemoveDownVoteOfAnAccount(reaction *DataSignatures.VoteReaction) error {
	db := reactionQuery.dbClient.GetDB()

	query, err := db.Prepare(`UPDATE reaction
									SET down_vote = reaction.down_vote - 1
									WHERE down_vote > 0 AND review_id = $1 AND account_id = $2`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	if err != nil {
		log.Fatal(err)
	}

	_, err = query.Exec(reaction.ReviewID, reaction.AccountID)

	if err != nil {
		return err
	}

	return nil
}
