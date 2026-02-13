package usecases

import (
	"go-api/domain/entities"
	"go-api/domain/repositories"
)

type GetAlbumByIDUsecase struct {
	Repo repositories.AlbumRepository
}

func NewGetAlbumByIdUsecase(repo repositories.AlbumRepository) *GetAlbumByIDUsecase {
	return &GetAlbumByIDUsecase{
		Repo: repo,
	}
}

func (uc *GetAlbumByIDUsecase) Execute(id string) (*entities.Album, error) {
	return uc.Repo.GetByID(id)
}
