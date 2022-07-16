package DataSignatures

type PostAddress struct {
	Country string `json:"country"`
	City    string `json:"city"`
	Street  string `json:"street"`
	Plaque  string `json:"plaque"`
}

type Address struct {
	Id uint64 `json:"id"`
	PostAddress
}
