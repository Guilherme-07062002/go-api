package factories

import (
	"go-api/controllers"
	inmemorydb "go-api/infra"
	"go-api/usecases"
)

func GetAlbumByIdFactory() *controllers.GetAlbumByIdController {
	repo := inmemorydb.NewAlbumRepository()
	usecase := usecases.NewGetAlbumByIdUsecase(repo)
	controller := controllers.NewGetAlbumByIDController(usecase)
	return controller
}
