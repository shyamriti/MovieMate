package models

import "time"

type Movie struct {
	Id          uint      `gorm:"primaryKey" json:"id"`
	Tile        string    `gorm:"not null" json:"title"`
	Year        int       `gorm:"not null" json:"year"`
	Genre       string    `gorm:"not null" json:"genre"`
	Description string    `gorm:"not null" json:"description"`
	Duration    int       `gorm:"not null" json:"duration"`
	ReleaseDate string    `gorm:"not null" json:"release_date"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}
