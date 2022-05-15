package repositories

import "github.com/cleysopnph/adote-um-pet/core/models"

type AdoptionRepository interface {
	FindAll() []models.Adoption
	Create(model models.Adoption) (models.Adoption, error)
	ExistsById(id uint) bool
}

type PetRepository interface {
	FindAll() []models.Pet
	Create(model models.Pet) (models.Pet, error)
	ExistsById(id uint) bool
}
