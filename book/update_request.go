package book

type BookUpdateRequest struct {
	Title       string `json:"title"`
	Price       int    `json:"price" binding:"number"`
	Description string `json:"description"`
	Rating      int    `json:"rating" binding:"number"`
	Discount    int    `json:"discount" binding:"number"`
}
