package controllers

import (
	"go-api/domain/dtos"
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

func (controller *UpdateAlbumController) Handle(c *gin.Context) {
	var body dtos.UpdateAlbumDto
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := controller.validator.Struct(body); err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "erro de validação",
			"details": err.Error(),
		})
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
