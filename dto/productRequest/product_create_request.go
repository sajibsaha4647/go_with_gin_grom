package productRequest

import "time"

type ProductCreateRequest struct {
	Title       string    `json:"title" gorm:"type:varchar(255);not null"`
	Description string    `json:"description" gorm:"type:text"`
	Price       float64   `json:"price" gorm:"type:numeric(10,2);not null;default:0.00"`
	ImageURL    string    `json:"imageUrl" gorm:"type:varchar(512);default:'default.png'"`
	Image       string    `json:"image" gorm:"type:varchar(255);default:'default.png'"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
