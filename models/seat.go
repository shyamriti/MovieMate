package models

import "time"

type Seats struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	SeatClass string    `gorm:"not null" json:"seat_class"`
	Price     int       `gorm:"not null" json:"price"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
}
