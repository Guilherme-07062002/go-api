package controllers

import (
	"go-api/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAlbumByIdController struct {
	GetAlbumByIDUsecase *usecases.GetAlbumByIDUsecase
}

func NewGetAlbumByIDController(usecase *usecases.GetAlbumByIDUsecase) *GetAlbumByIdController {
	return &GetAlbumByIdController{
		GetAlbumByIDUsecase: usecase,
	}
}

func (controller *GetAlbumByIdController) Handle(c *gin.Context) {
	id := c.Param("id")
	album, err := controller.GetAlbumByIDUsecase.Execute(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
	}

	c.IndentedJSON(http.StatusOK, album)
}
