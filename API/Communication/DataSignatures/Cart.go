package DataSignatures

type PostCart struct {
	ProductId       uint64    `json:"productId" binding:"required"`
	ProductCount    uint64    `json:"productCount" binding:"required"`
}