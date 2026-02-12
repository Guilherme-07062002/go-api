package inmemorydb

import (
	"errors"
	"fmt"
	"go-api/domain/dtos"
	entities "go-api/domain/entities"
)

type AlbumRepositoryMemory struct {
	albums []entities.Album
}

var AlbumsInMemory = []entities.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func NewAlbumRepository() *AlbumRepositoryMemory {
	return &AlbumRepositoryMemory{
		albums: AlbumsInMemory,
	}
}

func (r *AlbumRepositoryMemory) GetAll() (*[]entities.Album, error) {
	return &r.albums, nil
}

func (r *AlbumRepositoryMemory) GetByID(id string) (entities.Album, error) {
	for _, a := range r.albums {
		if a.ID == id {
			return a, nil
		}
	}
	return entities.Album{}, errors.New("album not found")
}

func (r *AlbumRepositoryMemory) Create(album dtos.CreateAlbumDto) entities.Album {
	var newAlbum entities.Album
	newId := fmt.Sprint(len(AlbumsInMemory) + 1)
	newAlbum.ID = newId
	newAlbum.Title = album.Title
	newAlbum.Artist = album.Artist
	newAlbum.Price = album.Price

	AlbumsInMemory = append(AlbumsInMemory, newAlbum)
	return newAlbum
}
