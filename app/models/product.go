package models

import "time"

type Product struct {
	ID            uint            `gorm:"primary_key" json:"id"`
	Name          string          `gorm:"varchar(255) not null" json:"name"`
	Slug          string          `gorm:"varchar(255) not null" json:"slug"`
	Type          string          `gorm:"varchar(255) not null" json:"type"`
	Description   string          `json:"description"`
	Price         int             `gorm:"not null" json:"price"`
	Quantity      int             `gorm:"not null" json:"quantity"`
	ProductImages []ProductImages `json:"attachments"`
	CreatedByID   uint            `gorm:"not null" json:"created_by"`
	CreatedBy     User            `gorm:"foreignKey(CreatedByID)" json:"-"`
	CreatedAt     time.Time       `json:"-"`
	UpdatedAt     time.Time       `json:"-"`
}

type ProductImages struct {
	ID         int       `gorm:"primary_key" json:"id"`
	ProductID  uint      `gorm:"not null" json:"product_id"`
	Url        string    `json:"url"`
	IsFeatured bool      `json:"is_featured"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
}
