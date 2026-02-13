package usecases

import (
	"go-api/domain/dtos"
	"go-api/domain/entities"
	"go-api/domain/repositories"
)

type UpdateAlbumUsecase struct {
	Repo repositories.AlbumRepository
}

func NewUpdateAlbumUsecase(repo repositories.AlbumRepository) *UpdateAlbumUsecase {
	return &UpdateAlbumUsecase{
		Repo: repo,
	}
}

func (uc *UpdateAlbumUsecase) Execute(id string, data dtos.UpdateAlbumDto) (*entities.Album, error) {
	return uc.Repo.Update(id, data)
}
