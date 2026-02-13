package controllers

import (
	"go-api/domain/dtos"
	"go-api/domain/exceptions"
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
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	result, err := controller.UpdateAlbumUsecase.Execute(id, body)
	if err != nil {
		if err == exceptions.AlbumNotFound {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}
