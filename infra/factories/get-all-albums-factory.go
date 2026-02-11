package factories

import (
	"go-api/controllers"
	inmemorydb "go-api/infra"
	"go-api/usecases"
)

func GetAllAlbumFactory() *controllers.GetAllAlbumsController {
	repo := inmemorydb.NewAlbumRepository()
	usecase := usecases.NewGetAlbumsUsecase(repo)
	controller := controllers.NewGetAllAlbumsController(usecase)
	return controller
}
