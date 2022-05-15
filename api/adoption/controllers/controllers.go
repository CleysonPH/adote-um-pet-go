package controllers

import "github.com/gin-gonic/gin"

type AdoptionController interface {
	Create(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}
