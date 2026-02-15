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

// GetAllAlbums godoc
// @Summary      Lista os álbuns cadastrados
// @Tags         Álbuns
// @Accept       json
// @Produce      json
// @Success      200  {array}  entities.Album "Lista de álbuns"
// @Router       /albums [get]
func (controller *GetAllAlbumsController) Handle(c *gin.Context) {
	albums, err := controller.GetAllUsecase.Execute()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, albums)
}
