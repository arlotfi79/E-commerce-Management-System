package DataSignatures

type ReviewWithVotes struct {
	Id            uint64 `json:"id"`
	Description   string `json:"description"`
	Rating        uint64 `json:"rating"`
	UpVoteCount   uint64 `json:"up_vote_count"`
	DownVoteCount uint64 `json:"down_vote_count"`
	RatingAVG     uint64 `json:"rating_avg"`
}

type PostReview struct {
	ProductID   uint64 `json:"product_id"`
	AccountID   uint64 `json:"account_id"`
	Description string `json:"description"`
	Rating      uint64 `json:"rating"`
}
