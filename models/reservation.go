package models

import "time"

type Reservation struct {
	Id              uint      `gorm:"primaryKey" json:"id"`
	MovieId         uint      `gorm:"not null" json:"movie_id"`
	UserId          uint      `gorm:"not null" json:"user_id"`
	Seats           int       `gorm:"not null" json:"seats"`
	ReservationDate time.Time `gorm:"not null" json:"reaservation_date"`
	CreatedAt       time.Time `gorm:"not null" json:"created_at"`
}
