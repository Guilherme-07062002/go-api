package factories

import (
	"go-api/controllers"
	"go-api/domain/repositories"
	"go-api/usecases"
)

func UpdateAlbumFactory(repo repositories.AlbumRepository) *controllers.UpdateAlbumController {
	usecase := usecases.NewUpdateAlbumUsecase(repo)
	controller := controllers.NewUpdateAlbumController(usecase)
	return controller
}
