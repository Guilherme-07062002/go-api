package controllers

import (
	"go-api/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAverageAlbumPricesController struct {
	GetAverageAlbumPricesUsecase *usecases.GetAverageAlbumPricesUsecase
}

func NewGetAverageAlbumsPriceController(usecase *usecases.GetAverageAlbumPricesUsecase) *GetAverageAlbumPricesController {
	return &GetAverageAlbumPricesController{
		GetAverageAlbumPricesUsecase: usecase,
	}
}

// GetAverageAlbumsPrice godoc
// @Summary      Retorna a média de preços dos álbuns cadastrados
// @Tags         Álbuns
// @Accept       json
// @Produce      json
// @Success      200  {object}  dtos.GetAverageAlbunsPriceResponseDto "Valor da média de preços dos álbuns"
// @Router       /albums/average [get]
// @Security     BearerAuth
func (controller *GetAverageAlbumPricesController) Handle(c *gin.Context) {
	response := controller.GetAverageAlbumPricesUsecase.Execute(c.Request.Context())
	c.IndentedJSON(http.StatusOK, response)
}
