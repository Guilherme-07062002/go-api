package controllers

import (
	"go-api/domain/dtos"
	"go-api/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateAlbumController struct {
	CreateAlbumUsecase *usecases.CreateAlbumUsecase
}

func NewCreateAlbumController(usecase *usecases.CreateAlbumUsecase) *CreateAlbumController {
	return &CreateAlbumController{
		CreateAlbumUsecase: usecase,
	}
}

func (controller *CreateAlbumController) Handle(c *gin.Context) {
	var body dtos.CreateAlbumDto
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})
		return
	}
	result := controller.CreateAlbumUsecase.Execute(body)
	c.IndentedJSON(http.StatusCreated, result)
}
