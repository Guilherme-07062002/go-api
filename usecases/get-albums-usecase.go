package usecases

import (
	"go-api/domain/entities"
	"go-api/domain/repositories"
)

type GetAlbumsUsecase struct {
	Repo repositories.AlbumRepository
}

func NewGetAlbumsUsecase(repo repositories.AlbumRepository) *GetAlbumsUsecase {
	return &GetAlbumsUsecase{
		Repo: repo,
	}
}

func (uc *GetAlbumsUsecase) Execute() (*[]entities.Album, error) {
	return uc.Repo.GetAll()
}
