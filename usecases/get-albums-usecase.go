package usecase

import (
	entities "go-api/domain/entities"
	repository "go-api/domain/repositories"
)

type GetAlbumsUsecase struct {
	Repo repository.AlbumRepository
}

func NewGetAlbumsUsecase(repo repository.AlbumRepository) *GetAlbumsUsecase {
	return &GetAlbumsUsecase{
		Repo: repo,
	}
}

func (uc *GetAlbumsUsecase) Execute() ([]entities.Album, error) {
	return uc.Repo.GetAll()
}
