package dtos

import "github.com/cleysopnph/adote-um-pet/api/pet/dtos"

type AdoptionResponse struct {
	ID    uint             `json:"id"`
	Email string           `json:"email"`
	Value float64          `json:"valor"`
	Pet   dtos.PetResponse `json:"pet"`
}

type AdoptionRequest struct {
	Email string  `json:"email" binding:"required,max=255,email"`
	Value float64 `json:"valor" binding:"required,min=10,max=100"`
	PetID uint    `json:"pet_id" binding:"required,min=1,petexists"`
}
