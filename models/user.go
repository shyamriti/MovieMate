package models

import (
	"time"
)

type User struct {
	Id          uint      `gorm:"primaryKey" json:"id"`
	UserName    string    `gorm:"not null" json:"user_name"`
	Password    string    `gorm:"not null" json:"password"`
	PhoneNumber int       `gorm:"not null" json:"phone_number"`
	Email       string    `gorm:"not null" json:"email"`
	Status      bool      `gorm:"not null" json:"status"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at"`
}
