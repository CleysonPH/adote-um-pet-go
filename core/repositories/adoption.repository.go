package repositories

import (
	"github.com/cleysopnph/adote-um-pet/core/models"
	"gorm.io/gorm"
)

func NewAdoptionRepository(connection *gorm.DB) AdoptionRepository {
	return &adoptionRepository{connection: connection}
}

type adoptionRepository struct {
	connection *gorm.DB
}

// ExistsById implements AdoptionRepository
func (r *adoptionRepository) ExistsById(id uint) bool {
	return r.connection.Find(&models.Adoption{}, id).RowsAffected > 0
}

// Create implements AdoptionRepository
func (r *adoptionRepository) Create(model models.Adoption) (models.Adoption, error) {
	if err := r.connection.First(&model.Pet, model.PetID).Error; err != nil {
		return models.Adoption{}, err
	}
	if err := r.connection.Create(&model).Error; err != nil {
		return models.Adoption{}, err
	}
	return model, nil
}

// FindAll implements AdoptionRepository
func (r *adoptionRepository) FindAll() []models.Adoption {
	var adoptions []models.Adoption
	r.connection.Model(&models.Adoption{}).Joins("Pet").Find(&adoptions)
	return adoptions
}
