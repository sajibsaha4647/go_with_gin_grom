package productRequest

type ProductCreateRequest struct {
	Title       string  `json:"title" form:"title" binding:"required"`
	Description string  `json:"description" form:"description"`
	Price       float64 `json:"price" form:"price" binding:"required,gt=0"`
	ImageURL    string  `json:"imageUrl"`
	Image       string  `json:"image"`
}
