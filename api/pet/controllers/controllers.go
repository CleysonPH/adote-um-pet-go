package controllers

import "github.com/gin-gonic/gin"

type PetController interface {
	Create(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}
