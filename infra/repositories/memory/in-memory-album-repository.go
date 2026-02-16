package memory

import (
	"context"
	dtos "go-api/domain/dtos/album"
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

func (r *InMemoryAlbumRepository) GetAll(ctx context.Context, page, limit int) (*[]entities.Album, int64, error) {
	total := int64(len(r.albums))
	offset := (page - 1) * limit

	// Verifica se o offset está fora do range
	if offset > len(r.albums) {
		emptyResult := make([]entities.Album, 0)
		return &emptyResult, total, nil
	}

	// Calcula o final da página
	end := offset + limit
	if end > len(r.albums) {
		end = len(r.albums)
	}

	paginatedAlbums := r.albums[offset:end]
	return &paginatedAlbums, total, nil
}

func (r *InMemoryAlbumRepository) GetByID(ctx context.Context, id string) (*entities.Album, error) {
	for _, a := range r.albums {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, exceptions.AlbumNotFound
}

func (r *InMemoryAlbumRepository) Create(ctx context.Context, album dtos.CreateAlbumDto) entities.Album {
	var newAlbum entities.Album
	newAlbum.ID = uuid.NewV4().String()
	newAlbum.Title = album.Title
	newAlbum.Artist = album.Artist
	newAlbum.Price = album.Price

	r.albums = append(r.albums, newAlbum)
	return newAlbum
}

func (r *InMemoryAlbumRepository) Update(ctx context.Context, id string, data dtos.UpdateAlbumDto) (*entities.Album, error) {
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
