package main

import (
	"go-api/domain/dtos"
	"go-api/infra/factories"
	"go-api/infra/middlewares"
	"go-api/infra/mocks"
	inmemorydb "go-api/infra/repositories"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	albums := mocks.GetAlbumsInMemory()

	repo := inmemorydb.NewAlbumRepository(albums)

	getAllAlbumController := factories.GetAllAlbumFactory(repo)
	router.GET("/albums", getAllAlbumController.Handle)

	getAlbumByIdController := factories.GetAlbumByIdFactory(repo)
	router.GET("/albums/:id", getAlbumByIdController.Handle)

	createAlbumController := factories.CreateAlbumFactory(repo)
	router.POST("/albums",
		middlewares.ValidateBody[dtos.CreateAlbumDto](),
		createAlbumController.Handle,
	)

	updateAlbumController := factories.UpdateAlbumFactory(repo)
	router.PUT("/albums/:id",
		middlewares.ValidateBody[dtos.UpdateAlbumDto](),
		updateAlbumController.Handle,
	)

	router.Run("localhost:8080")
}
