package memory

import (
	"go-api/domain/dtos"
	"go-api/domain/entities"
	"go-api/domain/exceptions"

	uuid "github.com/satori/go.uuid"
)

type InMemoryAlbumRepository struct {
	albums []entities.Album
}

func NewAlbumRepository(albums []entities.Album) *InMemoryAlbumRepository {
	return &InMemoryAlbumRepository{
		albums: albums,
	}
}

func (r *InMemoryAlbumRepository) GetAll() (*[]entities.Album, error) {
	return &r.albums, nil
}

func (r *InMemoryAlbumRepository) GetByID(id string) (*entities.Album, error) {
	for _, a := range r.albums {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, exceptions.AlbumNotFound
}

func (r *InMemoryAlbumRepository) Create(album dtos.CreateAlbumDto) entities.Album {
	var newAlbum entities.Album
	newAlbum.ID = uuid.NewV4().String()
	newAlbum.Title = album.Title
	newAlbum.Artist = album.Artist
	newAlbum.Price = album.Price

	r.albums = append(r.albums, newAlbum)
	return newAlbum
}

func (r *InMemoryAlbumRepository) Update(id string, data dtos.UpdateAlbumDto) (*entities.Album, error) {
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
