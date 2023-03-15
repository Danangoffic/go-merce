package models

import "time"

type UserToken struct {
	ID        uint      `gorm:"primary_key"`
	UserID    uint      `json:"user_id"`
	User      User      `gorm:"foreignKey(UserID)"`
	Token     string    `gorm:"not null" json:"token"`
	ExpiredAt time.Time `gorm:"not nill" json:"expired_at"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}
