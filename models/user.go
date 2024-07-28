package model

import "time"

type User struct {
	Id        uint `gorm:"primaryKey"`
	Name      *string
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
