package usecases

import (
	"context"
	dtos "go-api/domain/dtos/album"
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

func (uc *CreateAlbumUsecase) Execute(ctx context.Context, data dtos.CreateAlbumDto) entities.Album {
	return uc.Repo.Create(ctx, data)
}
