//go:build wireinject
// +build wireinject

package config

import (
	"go-api/controllers"
	"go-api/domain/dtos"
	"go-api/domain/repositories"
	"go-api/infra/middlewares"
	"go-api/infra/mocks"
	inmemorydb "go-api/infra/repositories"
	"go-api/usecases"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var albumRepositorySet = wire.NewSet(
	inmemorydb.NewAlbumRepository,
	mocks.GetAlbumsInMemory,
	wire.Bind(new(repositories.AlbumRepository), new(*inmemorydb.AlbumRepositoryMemory)),
)

var usecasesSet = wire.NewSet(
	usecases.NewGetAlbumsUsecase,
	usecases.NewCreateAlbumUsecase,
	usecases.NewGetAlbumByIdUsecase,
	usecases.NewUpdateAlbumUsecase,
)

var controllersSet = wire.NewSet(
	controllers.NewGetAllAlbumsController,
	controllers.NewCreateAlbumController,
	controllers.NewGetAlbumByIDController,
	controllers.NewUpdateAlbumController,
)

func newServer(
	getAllAlbumsController *controllers.GetAllAlbumsController,
	createAlbumController *controllers.CreateAlbumController,
	getAlbumByIdController *controllers.GetAlbumByIdController,
	updateAlbumController *controllers.UpdateAlbumController,
) *gin.Engine {
	router := gin.Default()

	router.GET("/albums", getAllAlbumsController.Handle)
	router.GET("/albums/:id", getAlbumByIdController.Handle)
	router.POST("/albums",
		middlewares.ValidateBody[dtos.CreateAlbumDto](),
		createAlbumController.Handle,
	)
	router.PUT("/albums/:id",
		middlewares.ValidateBody[dtos.UpdateAlbumDto](),
		updateAlbumController.Handle,
	)

	return router
}

func InitializeServer() *gin.Engine {
	wire.Build(
		albumRepositorySet,
		usecasesSet,
		controllersSet,
		newServer,
	)
	return nil
}
