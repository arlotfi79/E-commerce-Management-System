package DataSignatures

type ReviewWithVotes struct {
	Id            uint64 `json:"id"`
	Rating        uint64 `json:"rating"`
	UpVoteCount   uint64 `json:"up_vote_count"`
	DownVoteCount uint64 `json:"down_vote_count"`
}
