package factories

import (
	"go-api/controllers"
	"go-api/domain/repositories"
	"go-api/usecases"
)

func CreateAlbumFactory(repo repositories.AlbumRepository) *controllers.CreateAlbumController {
	usecase := usecases.NewCreateAlbumUsecase(repo)
	controller := controllers.NewCreateAlbumController(usecase)
	return controller
}
