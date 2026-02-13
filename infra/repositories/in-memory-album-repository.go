package inmemorydb

import (
	"fmt"
	"go-api/domain/dtos"
	"go-api/domain/entities"
	"go-api/domain/exceptions"
)

type AlbumRepositoryMemory struct {
	albums []entities.Album
}

func NewAlbumRepository(albums []entities.Album) *AlbumRepositoryMemory {
	return &AlbumRepositoryMemory{
		albums: albums,
	}
}

func (r *AlbumRepositoryMemory) GetAll() (*[]entities.Album, error) {
	return &r.albums, nil
}

func (r *AlbumRepositoryMemory) GetByID(id string) (*entities.Album, error) {
	for _, a := range r.albums {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, exceptions.AlbumNotFound
}

func (r *AlbumRepositoryMemory) Create(album dtos.CreateAlbumDto) entities.Album {
	var newAlbum entities.Album
	newId := fmt.Sprint(len(r.albums) + 1)
	newAlbum.ID = newId
	newAlbum.Title = album.Title
	newAlbum.Artist = album.Artist
	newAlbum.Price = album.Price

	r.albums = append(r.albums, newAlbum)
	return newAlbum
}

func (r *AlbumRepositoryMemory) Update(id string, data dtos.UpdateAlbumDto) (*entities.Album, error) {
	for i, a := range r.albums {
		if a.ID == id {
			albumFound := &r.albums[i]
			if data.Title != nil {
				albumFound.Title = *data.Title
			}
			if data.Price != nil {
				albumFound.Price = *data.Price
			}
			if data.Artist != nil {
				albumFound.Artist = *data.Artist
			}
			return albumFound, nil
		}
	}

	return nil, exceptions.AlbumNotFound
}
