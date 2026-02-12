package factories

import (
	"go-api/controllers"
	inmemorydb "go-api/infra/repositories"
	"go-api/usecases"
)

func CreateAlbumFactory() *controllers.CreateAlbumController {
	repo := inmemorydb.NewAlbumRepository()
	usecase := usecases.NewCreateAlbumUsecase(repo)
	controller := controllers.NewCreateAlbumController(usecase)
	return controller
}
