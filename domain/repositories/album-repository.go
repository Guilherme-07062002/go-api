package repositories

import (
	"go-api/domain/dtos"
	entities "go-api/domain/entities"
)

type AlbumRepository interface {
	GetByID(id string) (entities.Album, error)
	GetAll() (*[]entities.Album, error)
	Create(album dtos.CreateAlbumDto) entities.Album
}
