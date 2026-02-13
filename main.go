package main

import (
	"go-api/infra/factories"
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
	router.POST("/albums", createAlbumController.Handle)

	updateAlbumController := factories.UpdateAlbumFactory(repo)
	router.PUT("/albums/:id", updateAlbumController.Handle)

	router.Run("localhost:8080")
}
