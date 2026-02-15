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

// CreateAlbum godoc
// @Summary      Cria um novo álbum
// @Tags         Álbuns
// @Param        album  body      dtos.CreateAlbumDto  true "Dados do álbum que será criado"
// @Success      201    {object}  entities.Album "Álbum criado com sucesso"
// @Router       /albums [post]
// @Security BearerAuth
func (controller *CreateAlbumController) Handle(c *gin.Context) {
	val, _ := c.Get("validatedBody")
	dto := val.(dtos.CreateAlbumDto)

	result := controller.CreateAlbumUsecase.Execute(dto)
	c.IndentedJSON(http.StatusCreated, result)
}
