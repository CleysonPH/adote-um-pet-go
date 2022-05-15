package models

type Pet struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"size:100;not null"`
	History string `gorm:"size:255;not null"`
	Picture string `gorm:"size:255;not null"`
}
