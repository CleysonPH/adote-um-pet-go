package controllers

import (
	"errors"
	"net/http"

	"github.com/cleysopnph/adote-um-pet/api/adoption/dtos"
	"github.com/cleysopnph/adote-um-pet/api/adoption/mappers"
	"github.com/cleysopnph/adote-um-pet/core/repositories"
	"github.com/cleysopnph/adote-um-pet/core/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func NewAdoptionController(
	adoptionMapper mappers.AdoptionMapper,
	adoptionRepository repositories.AdoptionRepository,
) AdoptionController {
	return &adoptionController{
		adoptionMapper:     adoptionMapper,
		adoptionRepository: adoptionRepository,
	}
}

type adoptionController struct {
	adoptionRepository repositories.AdoptionRepository
	adoptionMapper     mappers.AdoptionMapper
}

// Create implements AdoptionController
func (c *adoptionController) Create(ctx *gin.Context) {
	var request dtos.AdoptionRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{
					"message": "Houveram erros de validação",
					"errors":  utils.GetValidationErrors(verr),
				},
			)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	adoption, err := c.adoptionRepository.Create(c.adoptionMapper.ToModel(request))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, c.adoptionMapper.ToResponse(adoption))
}

// FindAll implements AdoptionController
func (c *adoptionController) FindAll(ctx *gin.Context) {
	adoptions := c.adoptionRepository.FindAll()
	ctx.JSON(http.StatusOK, c.adoptionMapper.ToCollectionResponse(adoptions))
}
