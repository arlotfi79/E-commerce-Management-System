package DataSignatures

type Notification struct {
	Id          uint64 `json:"id"`
	AccountID   uint64 `json:"account_id"`
	Description string `json:"description"`
}
