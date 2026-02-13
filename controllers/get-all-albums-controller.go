package controllers

import (
	"go-api/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllAlbumsController struct {
	GetAllUsecase *usecases.GetAlbumsUsecase
}

func NewGetAllAlbumsController(usecase *usecases.GetAlbumsUsecase) *GetAllAlbumsController {
	return &GetAllAlbumsController{
		GetAllUsecase: usecase,
	}
}

func (controller *GetAllAlbumsController) Handle(c *gin.Context) {
	albums, err := controller.GetAllUsecase.Execute()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	c.IndentedJSON(http.StatusOK, albums)
}
