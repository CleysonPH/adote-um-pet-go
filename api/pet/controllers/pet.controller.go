package controllers

import (
	"errors"
	"net/http"

	"github.com/cleysopnph/adote-um-pet/api/pet/dtos"
	"github.com/cleysopnph/adote-um-pet/api/pet/mappers"
	"github.com/cleysopnph/adote-um-pet/core/repositories"
	"github.com/cleysopnph/adote-um-pet/core/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func NewPetController(
	petRepository repositories.PetRepository,
	petMapper mappers.PetMapper,
) PetController {
	return &petController{
		petRepository: petRepository,
		petMapper:     petMapper,
	}
}

type petController struct {
	petRepository repositories.PetRepository
	petMapper     mappers.PetMapper
}

// Create implements PetController
func (c *petController) Create(ctx *gin.Context) {
	var request dtos.PetRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{
					"message": "houveram erros de validação",
					"errors":  utils.GetValidationErrors(verr),
				},
			)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	pet, err := c.petRepository.Create(c.petMapper.ToModel(request))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, c.petMapper.ToResponse(pet))
}

// FindAll implements PetController
func (c *petController) FindAll(ctx *gin.Context) {
	pets := c.petRepository.FindAll()
	response := c.petMapper.ToCollectionResponse(pets)
	ctx.JSON(http.StatusOK, response)
}
