package repositories

import (
	"github.com/cleysopnph/adote-um-pet/core/models"
	"gorm.io/gorm"
)

func NewPetRepository(connection *gorm.DB) PetRepository {
	return &petRepository{connection: connection}
}

type petRepository struct {
	connection *gorm.DB
}

// ExistsById implements PetRepository
func (r *petRepository) ExistsById(id uint) bool {
	return r.connection.Find(&models.Pet{}, id).RowsAffected > 0
}

// Create implements PetRepository
func (r *petRepository) Create(model models.Pet) (models.Pet, error) {
	if err := r.connection.Create(&model).Error; err != nil {
		return models.Pet{}, err
	}
	return model, nil
}

// FindAll implements PetRepository
func (r *petRepository) FindAll() []models.Pet {
	var pets []models.Pet
	r.connection.Find(&pets)
	return pets
}
