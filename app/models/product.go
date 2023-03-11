package models

import "time"

type Product struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	CreatedByID uint      `gorm:"not null" json:"created_by_id"`
	CreatedBy   User      `gorm:"foreignKey(CreatedByID)" json:"created_by"`
	Name        string    `gorm:"varchar(255) not null" json:"name"`
	Slug        string    `gorm:"varchar(255) not null" json:"slug"`
	Type        string    `gorm:"varchar(255) not null" json:"type"`
	Description string    `json:"description"`
	Price       int       `gorm:"not null" json:"price"`
	Quantity    int       `gorm:"not null" json:"quantity"`
	CreatedAt   time.Time `json:"creatd_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
