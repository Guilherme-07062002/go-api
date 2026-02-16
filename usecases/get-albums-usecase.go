package usecases

import (
	"context"
	dtos "go-api/domain/dtos/pagination"
	"go-api/domain/entities"
	"go-api/domain/repositories"
	"math"
)

type GetAlbumsUsecase struct {
	Repo repositories.AlbumRepository
}

func NewGetAlbumsUsecase(repo repositories.AlbumRepository) *GetAlbumsUsecase {
	return &GetAlbumsUsecase{
		Repo: repo,
	}
}

func (uc *GetAlbumsUsecase) Execute(ctx context.Context, page, limit int) (*dtos.PaginatedResponse[entities.Album], error) {
	albums, total, err := uc.Repo.GetAll(ctx, page, limit)
	if err != nil {
		return nil, err
	}

	totalPages := int64(math.Ceil(float64(total) / float64(limit)))

	response := &dtos.PaginatedResponse[entities.Album]{
		Data:       *albums,
		Total:      total,
		Page:       int64(page),
		Limit:      int64(limit),
		TotalPages: totalPages,
	}

	return response, nil
}
