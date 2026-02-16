//go:build wireinject

package wire

import (
	"go-api/controllers"
	"go-api/domain/dtos"
	"go-api/domain/repositories"
	"go-api/infra/config/env"
	"go-api/infra/config/postgres"
	"go-api/infra/middlewares"
	"go-api/infra/mocks"
	"go-api/infra/repositories/memory"
	postgresRepository "go-api/infra/repositories/postgres"
	"go-api/infra/security"

	"go-api/usecases"

	_ "go-api/docs"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func provideJWTSecret() string {
	env := env.LoadEnv()
	return env.JwtSecret
}

var securitySet = wire.NewSet(
	provideJWTSecret,
	security.NewJwtService,
)

var albumRepositorySet = wire.NewSet(
	memory.NewAlbumRepository,
	mocks.GetAlbumsInMemory,
	wire.Bind(new(repositories.AlbumRepository), new(*memory.InMemoryAlbumRepository)),
)

func provideDbInstance() *gorm.DB {
	DB := postgres.DB
	return DB
}

var postgresRepositorySet = wire.NewSet(
	postgresRepository.NewPostgresRepository,
	provideDbInstance,
	wire.Bind(new(repositories.AlbumRepository), new(*postgresRepository.PostgresAlbumRepository)),
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
	tokenService security.TokenService,
) *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// router.Use(middlewares.AuthMiddleware(tokenService))

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
		postgresRepositorySet,
		usecasesSet,
		controllersSet,
		securitySet,
		newServer,
	)
	return nil
}
