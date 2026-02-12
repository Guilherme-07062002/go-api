package usecases

import (
	"go-api/domain/dtos"
	"go-api/domain/entities"
	"go-api/domain/repositories"
)

type CreateAlbumUsecase struct {
	Repo repositories.AlbumRepository
}

func NewCreateAlbumUsecase(repo repositories.AlbumRepository) *CreateAlbumUsecase {
	return &CreateAlbumUsecase{
		Repo: repo,
	}
}

func (uc *CreateAlbumUsecase) Execute(data dtos.CreateAlbumDto) entities.Album {
	return uc.Repo.Create(data)
}
