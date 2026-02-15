package controllers

import (
	"go-api/domain/exceptions"
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

// GetAlbumByID godoc
// @Summary      Busca um álbum pelo seu ID
// @Tags         Álbuns
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "ID do álbum que será buscado"
// @Success      200  {object}  entities.Album "Álbum encontrado"
// @Router       /albums/{id} [get]
// @Security BearerAuth
func (controller *GetAlbumByIdController) Handle(c *gin.Context) {
	id := c.Param("id")
	album, err := controller.GetAlbumByIDUsecase.Execute(id)
	if err != nil {
		if err == exceptions.AlbumNotFound {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}

		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}
