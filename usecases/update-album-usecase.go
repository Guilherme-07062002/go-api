package usecases

import (
	"context"
	dtos "go-api/domain/dtos/album"
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

func (uc *UpdateAlbumUsecase) Execute(ctx context.Context, id string, data dtos.UpdateAlbumDto) (*entities.Album, error) {
	return uc.Repo.Update(ctx, id, data)
}
