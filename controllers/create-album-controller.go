package controllers

import (
	"go-api/domain/dtos"
	"go-api/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateAlbumController struct {
	CreateAlbumUsecase *usecases.CreateAlbumUsecase
	validator          *validator.Validate
}

func NewCreateAlbumController(usecase *usecases.CreateAlbumUsecase) *CreateAlbumController {
	return &CreateAlbumController{
		CreateAlbumUsecase: usecase,
		validator:          validator.New(),
	}
}

func (controller *CreateAlbumController) Handle(c *gin.Context) {
	var body dtos.CreateAlbumDto
	if err := c.ShouldBindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := controller.validator.Struct(&body); err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "erro de validação",
			"details": err.Error(),
		})
		return
	}
	result := controller.CreateAlbumUsecase.Execute(body)
	c.IndentedJSON(http.StatusCreated, result)
}
