package repositories

import (
	"context"
	dtos "go-api/domain/dtos/album"
	"go-api/domain/entities"
)

type AlbumRepository interface {
	GetByID(ctx context.Context, id string) (*entities.Album, error)
	GetAll(ctx context.Context, page, limit int) (*[]entities.Album, int64, error)
	Create(ctx context.Context, album dtos.CreateAlbumDto) entities.Album
	Update(ctx context.Context, id string, album dtos.UpdateAlbumDto) (*entities.Album, error)
}
