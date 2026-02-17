package mocks

import (
	"context"
	dtos "go-api/domain/dtos/album"
	"go-api/domain/entities"
)

type MockAlbumRepository struct {
	MockGetAllWithoutPagination func(ctx context.Context) *[]entities.Album
	MockGetByID                 func(ctx context.Context) (*entities.Album, error)
	MockGetAll                  func(ctx context.Context, page, limit int) (*[]entities.Album, int64, error)
	MockCreate                  func(ctx context.Context, album dtos.CreateAlbumDto) entities.Album
	MockUpdate                  func(ctx context.Context, id string, albums dtos.UpdateAlbumDto) (*entities.Album, error)
}

func (m *MockAlbumRepository) GetAllWithoutPagination(ctx context.Context) *[]entities.Album {
	return m.MockGetAllWithoutPagination(ctx)
}

func (m *MockAlbumRepository) GetByID(ctx context.Context, id string) (*entities.Album, error) {
	panic("Método não implementado")
}

func (m *MockAlbumRepository) GetAll(ctx context.Context, page, limit int) (*[]entities.Album, int64, error) {
	panic("Método não implementado")
}

func (m *MockAlbumRepository) Create(ctx context.Context, album dtos.CreateAlbumDto) entities.Album {
	panic("Método não implementado")
}

func (m *MockAlbumRepository) Update(ctx context.Context, id string, album dtos.UpdateAlbumDto) (*entities.Album, error) {
	panic("Método não implementado")
}
