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

func (controller *CreateAlbumController) Handle(c *gin.Context) {
	val, _ := c.Get("validatedBody")
	dto := val.(dtos.CreateAlbumDto)

	result := controller.CreateAlbumUsecase.Execute(dto)
	c.IndentedJSON(http.StatusCreated, result)
}
