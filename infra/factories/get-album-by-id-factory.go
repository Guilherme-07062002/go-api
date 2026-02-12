package factories

import (
	"go-api/controllers"
	"go-api/domain/repositories"
	"go-api/usecases"
)

func GetAlbumByIdFactory(repo repositories.AlbumRepository) *controllers.GetAlbumByIdController {
	usecase := usecases.NewGetAlbumByIdUsecase(repo)
	controller := controllers.NewGetAlbumByIDController(usecase)
	return controller
}
