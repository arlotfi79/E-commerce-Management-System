package Queries

import (
	"API/Communication/DataSignatures"
	"API/Database"
	"log"
)

type ReviewQuery struct {
	dbClient *Database.Postgresql
}

func NewReviewQuery(dbClient *Database.Postgresql) *ReviewQuery {
	return &ReviewQuery{dbClient: dbClient}
}

func (reviewQuery *ReviewQuery) GetReviewsWithVotesByProductID(id uint64) ([]DataSignatures.ReviewWithVotes, error) {
	db := reviewQuery.dbClient.GetDB()

	query, err := db.Prepare(`SELECT r.review_id, r.rating, SUM(re.up_vote), SUM(re.down_vote)
									FROM review AS r 
									INNER JOIN reaction AS re ON r.review_id = re.review_id
									WHERE r.product_id = $1
									GROUP BY r.review_id, r.rating`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	row, err := query.Query(id)

	if err != nil {
		log.Fatal(err)
	}

	var reviews []DataSignatures.ReviewWithVotes
	for row.Next() {
		var review DataSignatures.ReviewWithVotes
		row.Scan(&review.Id, &review.Rating, &review.UpVoteCount, &review.DownVoteCount)

		if err != nil {
			log.Fatal(err)
		}

		reviews = append(reviews, review)
	}

	return reviews, err
}
