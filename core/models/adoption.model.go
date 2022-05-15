package models

type Adoption struct {
	ID    uint    `gorm:"primaryKey"`
	Value float64 `gorm:"precision:5;scale:2;not null"`
	Email string  `gorm:"size:255;not null"`
	PetID uint    `gorm:"not null"`
	Pet   Pet
}
