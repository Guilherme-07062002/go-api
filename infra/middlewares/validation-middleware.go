package middlewares

import (
	"go-api/infra/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateBody[T any]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto T

		if err := c.ShouldBindJSON(&dto); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "JSON inválido",
			})
			return
		}

		if err := validate.Struct(dto); err != nil {
			errors := utils.TranslateError(err)
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"message": "Dados inválidos",
				"errors":  errors,
			})
			return
		}

		c.Set("validatedBody", dto)
		c.Next()
	}
}
