package models

import "time"

type Movie struct {
	Id          uint      `gorm:"primaryKey" json:"id"`
	Tile        string    `gorm:"not null" json:"title"`
	Year        int       `gorm:"not null" json:"year"`
	Genre       string    `gorm:"not null" json:"genre"`
	Description string    `gorm:"not null" json:"description"`
	Duration    int       `gorm:"not null" json:"duration"`
	ReleaseDate time.Time `gorm:"not null" json:"release_date"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at"`
}
