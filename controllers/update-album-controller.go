package controllers

import (
	"go-api/domain/dtos"
	"go-api/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateAlbumController struct {
	UpdateAlbumUsecase *usecases.UpdateAlbumUsecase
}

func NewUpdateAlbumController(usecase *usecases.UpdateAlbumUsecase) *UpdateAlbumController {
	return &UpdateAlbumController{
		UpdateAlbumUsecase: usecase,
	}
}

func (controller *UpdateAlbumController) Handle(c *gin.Context) {
	var body dtos.UpdateAlbumDto
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})
		return
	}
	id := c.Param("id")
	result, err := controller.UpdateAlbumUsecase.Execute(id, body)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}
