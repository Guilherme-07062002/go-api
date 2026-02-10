package repositories

import album "go-api/domain/entities"

type AlbumRepository interface {
	GetByID(id string) (album.Album, error)
	GetAll() ([]album.Album, error)
}
