package main

import (
	adoption_controller "github.com/cleysopnph/adote-um-pet/api/adoption/controllers"
	adoption_mapper "github.com/cleysopnph/adote-um-pet/api/adoption/mappers"
	pet_controller "github.com/cleysopnph/adote-um-pet/api/pet/controllers"
	pet_mapper "github.com/cleysopnph/adote-um-pet/api/pet/mappers"
	"github.com/cleysopnph/adote-um-pet/core/config"
	"github.com/cleysopnph/adote-um-pet/core/database"
	"github.com/cleysopnph/adote-um-pet/core/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = database.ConfigDatabase()

	petMapper     pet_mapper.PetMapper         = pet_mapper.NewPetMapper()
	petRepository repositories.PetRepository   = repositories.NewPetRepository(db)
	petController pet_controller.PetController = pet_controller.NewPetController(petRepository, petMapper)

	adoptionMapper     adoption_mapper.AdoptionMapper         = adoption_mapper.NewAdoptionMapper(petMapper)
	adoptionRepository repositories.AdoptionRepository        = repositories.NewAdoptionRepository(db)
	adoptionController adoption_controller.AdoptionController = adoption_controller.NewAdoptionController(adoptionMapper, adoptionRepository)
)

func main() {
	defer database.CloseDatabaseConnection(db)

	config.ConfigValidator(petRepository)

	r := gin.Default()

	apiRoutes := r.Group("/api")
	{
		petRoutes := apiRoutes.Group("/pets")
		{
			petRoutes.GET("/", petController.FindAll)
			petRoutes.POST("/", petController.Create)
		}
		adoptionRoutes := apiRoutes.Group("/adocoes")
		{
			adoptionRoutes.GET("/", adoptionController.FindAll)
			adoptionRoutes.POST("/", adoptionController.Create)
		}
	}

	r.Run(":8000")
}
