package controllers

import (
	dtos "go-api/domain/dtos/pagination"
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
// @Summary      Lista os álbuns cadastrados com paginação
// @Tags         Álbuns
// @Accept       json
// @Produce      json
// @Param        page   query     int  false  "Número da página" default(1) minimum(1)
// @Param        limit  query     int  false  "Itens por página" default(10) minimum(1) maximum(100)
// @Success      200  {object}  dtos.PaginatedResponse[entities.Album] "Lista paginada de álbuns"
// @Router       /albums [get]
// @Security     BearerAuth
func (controller *GetAllAlbumsController) Handle(c *gin.Context) {
	var paginationParams dtos.PaginationDto

	if err := c.ShouldBindQuery(&paginationParams); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Valores padrão caso não sejam fornecidos
	if paginationParams.Page == 0 {
		paginationParams.Page = 1
	}
	if paginationParams.Limit == 0 {
		paginationParams.Limit = 10
	}

	response, err := controller.GetAllUsecase.Execute(c.Request.Context(), paginationParams.Page, paginationParams.Limit)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, response)
}
