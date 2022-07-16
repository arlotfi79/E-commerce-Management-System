package DataSignatures

type PostReaction struct {
	ReviewId uint64 `json:"reviewId" binding:"required"`
}