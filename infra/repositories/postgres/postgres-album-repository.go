package postgres

import (
	"context"
	dtos "go-api/domain/dtos/album"
	"go-api/domain/entities"
	"go-api/domain/exceptions"
	postgresConfig "go-api/infra/config/postgres"
	"go-api/infra/repositories/postgres/models"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type PostgresAlbumRepository struct {
	DB *gorm.DB
}

func NewPostgresRepository() *PostgresAlbumRepository {
	return &PostgresAlbumRepository{DB: postgresConfig.DB}
}

func (r *PostgresAlbumRepository) GetAll(ctx context.Context, page, limit int) (*[]entities.Album, int64, error) {
	var albumModels []models.Album
	var total int64

	r.DB.WithContext(ctx).Model(&models.Album{}).Count(&total)

	offset := (page - 1) * limit
	result := r.DB.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Find(&albumModels)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	albums := make([]entities.Album, len(albumModels))
	for i, model := range albumModels {
		albums[i] = entities.Album{
			ID:     model.ID,
			Title:  model.Title,
			Artist: model.Artist,
			Price:  model.Price,
		}
	}

	return &albums, total, nil
}

func (r *PostgresAlbumRepository) GetByID(ctx context.Context, id string) (*entities.Album, error) {
	var albumModel models.Album
	result := r.DB.First(&albumModel, "id = ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, exceptions.AlbumNotFound
		}
		return nil, result.Error
	}

	album := &entities.Album{
		ID:     albumModel.ID,
		Title:  albumModel.Title,
		Artist: albumModel.Artist,
		Price:  albumModel.Price,
	}

	return album, nil
}

func (r *PostgresAlbumRepository) Create(ctx context.Context, data dtos.CreateAlbumDto) entities.Album {
	albumModel := models.Album{
		ID:     uuid.NewV4().String(),
		Title:  data.Title,
		Artist: data.Artist,
		Price:  data.Price,
	}

	r.DB.Create(&albumModel)

	return entities.Album{
		ID:     albumModel.ID,
		Title:  albumModel.Title,
		Artist: albumModel.Artist,
		Price:  albumModel.Price,
	}
}

func (r *PostgresAlbumRepository) Update(ctx context.Context, id string, data dtos.UpdateAlbumDto) (*entities.Album, error) {
	var albumModel models.Album
	result := r.DB.First(&albumModel, "id = ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, exceptions.AlbumNotFound
		}
		return nil, result.Error
	}

	if data.Title != nil {
		albumModel.Title = *data.Title
	}
	if data.Artist != nil {
		albumModel.Artist = *data.Artist
	}
	if data.Price != nil {
		albumModel.Price = *data.Price
	}

	r.DB.Save(&albumModel)

	album := &entities.Album{
		ID:     albumModel.ID,
		Title:  albumModel.Title,
		Artist: albumModel.Artist,
		Price:  albumModel.Price,
	}

	return album, nil
}
