package DataSignatures

type Reaction struct {
	Id       uint64 `json:"id"`
	UpVote   uint64 `json:"up_vote"`
	DownVote uint64 `json:"down_vote"`
}

type VoteReaction struct {
	ReviewID  uint64 `json:"review_id"`
	AccountID uint64 `json:"account_id"`
	Vote      uint64 `json:"vote"`
}
