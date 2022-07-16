package DataSignatures

type Category struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

type GetCategory struct {
	Name string `json:"name"`
}
