package usecases

import (
	"context"
	"go-api/domain/repositories"
)

type GetAverageAlbumPricesUsecase struct {
	Repo repositories.AlbumRepository
}

func NewGetAverageAlbumPricesUsecase(repo repositories.AlbumRepository) *GetAverageAlbumPricesUsecase {
	return &GetAverageAlbumPricesUsecase{
		Repo: repo,
	}
}

func (uc *GetAverageAlbumPricesUsecase) Execute(ctx context.Context) float64 {
	albumsPtr := uc.Repo.GetAllWithoutPagination(ctx)
	if len(*albumsPtr) == 0 || albumsPtr == nil {
		return 0
	}

	albums := *albumsPtr
	var totalSum int64 = 0

	for _, a := range albums {
		totalSum += int64(a.Price)
	}

	average := float64(totalSum) / float64(len(albums))
	return average
}
