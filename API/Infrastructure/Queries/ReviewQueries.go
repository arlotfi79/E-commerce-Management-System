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

// -------------------------------- GET ----------------------------------

func (reviewQuery *ReviewQuery) GetReviewsWithVotesByProductID(id uint64) ([]DataSignatures.ReviewWithVotes, error) {
	db := reviewQuery.dbClient.GetDB()

	query, err := db.Prepare(`SELECT r.review_id, r.rating, r.description,
									   SUM(CASE WHEN re.up_vote THEN 1 ELSE 0 END) AS up_vote_count,
									   SUM(CASE WHEN re.down_vote THEN 1 ELSE 0 END) AS down_vote_count,
									   avg_rating
									FROM review AS r
									LEFT JOIN reaction AS re ON r.review_id = re.review_id
									CROSS JOIN (
												SELECT AVG(rating) AS avg_rating
												FROM Review
												WHERE product_id = $1
												) AS rating_avg
									WHERE r.product_id = $1
									GROUP BY r.review_id, r.rating, r.description, avg_rating`)

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
		row.Scan(&review.Id, &review.Rating, &review.Description, &review.UpVoteCount, &review.DownVoteCount, &review.RatingAVG)

		if err != nil {
			log.Fatal(err)
		}

		reviews = append(reviews, review)
	}

	return reviews, err
}

// -------------------------------- POST ----------------------------------

func (reviewQuery *ReviewQuery) PostReviewUsingProductID(review *DataSignatures.PostReview) error {
	db := reviewQuery.dbClient.GetDB()

	query, err := db.Prepare(`INSERT INTO review (product_id, account_id, description, rating)
									VALUES ($1, $2, $3, $4)`)
	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	_, err = query.Exec(review.ProductID, review.AccountID, review.Description, review.Rating)

	if err != nil {
		return err
	}

	return nil
}
