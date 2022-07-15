package DataSignatures

type GetReply struct {
	Id        uint64 `json:"id"`
	ReplyText string `json:"reply_text"`
}

type PostReply struct {
	ReviewID  uint64 `json:"review_id"`
	ReplyText string `json:"reply_text"`
}
