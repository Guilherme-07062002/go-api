package repositories

import entities "go-api/domain/entities"

type AlbumRepository interface {
	GetByID(id string) (entities.Album, error)
	GetAll() ([]entities.Album, error)
}
