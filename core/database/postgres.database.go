package database

import (
	"log"

	"github.com/cleysopnph/adote-um-pet/core/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConfigDatabase() *gorm.DB {
	connectionString := "host=localhost user=root password=root dbname=root port=5432"
	db, err := gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados: " + err.Error())
	}
	db.AutoMigrate(&models.Pet{}, &models.Adoption{})
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		log.Panic("Erro ao fechar a conexão com o banco de dados: " + err.Error())
	}
	dbSQL.Close()
}
