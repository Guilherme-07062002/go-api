package repositories

import (
	"go-api/domain/dtos"
	"go-api/domain/entities"
)

type AlbumRepository interface {
	GetByID(id string) (entities.Album, error)
	GetAll() (*[]entities.Album, error)
	Create(album dtos.CreateAlbumDto) entities.Album
	Update(id string, album dtos.UpdateAlbumDto) (*entities.Album, error)
}
