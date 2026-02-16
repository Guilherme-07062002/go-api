package postgres

import (
	"go-api/domain/dtos"
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

func (r *PostgresAlbumRepository) GetAll() (*[]entities.Album, error) {
	var albumModels []models.Album
	result := r.DB.Find(&albumModels)
	if result.Error != nil {
		return nil, result.Error
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

	return &albums, nil
}

func (r *PostgresAlbumRepository) GetByID(id string) (*entities.Album, error) {
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

func (r *PostgresAlbumRepository) Create(data dtos.CreateAlbumDto) entities.Album {
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

func (r *PostgresAlbumRepository) Update(id string, data dtos.UpdateAlbumDto) (*entities.Album, error) {
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
