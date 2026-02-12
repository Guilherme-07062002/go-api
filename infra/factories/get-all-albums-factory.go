package factories

import (
	"go-api/controllers"
	"go-api/domain/repositories"
	"go-api/usecases"
)

func GetAllAlbumFactory(repo repositories.AlbumRepository) *controllers.GetAllAlbumsController {
	usecase := usecases.NewGetAlbumsUsecase(repo)
	controller := controllers.NewGetAllAlbumsController(usecase)
	return controller
}
