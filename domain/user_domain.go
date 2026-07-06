package domain

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Image	 string    `json:"image" gorm:"type:varchar(255);default:'default.png'"`
	Name      string    `json:"name" gorm:"type:varchar(100);not null"`
	Email     string    `json:"email" gorm:"type:varchar(255);unique;not null;index"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null"`
	UserType  string    `json:"userType" gorm:"type:varchar(50);default:'customer';column:user_type"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
