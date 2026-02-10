package inmemorydb

import (
	"errors"
	entities "go-api/domain/entities"
)

type AlbumRepositoryMemory struct {
	albums []entities.Album
}

func NewAlbumRepository() *AlbumRepositoryMemory {
	return &AlbumRepositoryMemory{
		albums: []entities.Album{
			{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
			{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
			{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
		},
	}
}

func (r *AlbumRepositoryMemory) GetAll() ([]entities.Album, error) {
	return r.albums, nil
}

func (r *AlbumRepositoryMemory) GetByID(id string) (entities.Album, error) {
	for _, a := range r.albums {
		if a.ID == id {
			return a, nil
		}
	}
	return entities.Album{}, errors.New("album not found")
}
