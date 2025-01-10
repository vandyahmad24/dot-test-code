package request

type BookRequest struct {
	Title      string `json:"title" binding:"required"`
	Author     string `json:"author" binding:"required"`
	CategoryID uint   `json:"category_id" binding:"required"`
}
