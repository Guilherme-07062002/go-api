package controllers

import (
	dtos "go-api/domain/dtos/album"
	"go-api/domain/exceptions"
	"go-api/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UpdateAlbumController struct {
	UpdateAlbumUsecase *usecases.UpdateAlbumUsecase
	validator          *validator.Validate
}

func NewUpdateAlbumController(usecase *usecases.UpdateAlbumUsecase) *UpdateAlbumController {
	return &UpdateAlbumController{
		UpdateAlbumUsecase: usecase,
		validator:          validator.New(),
	}
}

// UpdateAlbum godoc
// @Summary      Atualiza um álbum
// @Tags         Álbuns
// @Accept       json
// @Produce      json
// @Param        id     path      string             true  "ID do álbum que será atualizado"
// @Param        album  body      dtos.UpdateAlbumDto  true  "Informações do álbum que serão atualizadas"
// @Success      200    {object}  entities.Album "Álbum atualizado com sucesso"
// @Router       /albums/{id} [put]
// @Security BearerAuth
func (controller *UpdateAlbumController) Handle(c *gin.Context) {
	val, _ := c.Get("validatedBody")
	dto := val.(dtos.UpdateAlbumDto)

	id := c.Param("id")
	result, err := controller.UpdateAlbumUsecase.Execute(c.Request.Context(), id, dto)
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
