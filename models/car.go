package models

import (
	"time"

	"gorm.io/gorm"
)

type Car struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Model     string         `json:"model"`
	Color     string         `json:"color"`
	Brand     string         `json:"brand"`
	Year      string         `json:"year"`
	Motor     string         `json:"motor"`
	Power     string         `json:"power"`
	Traction  string         `json:"traction"`
	ImageURL  string         `json:"img_url"`
	CreatedAt time.Time      `json:"created"`
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}
