package usecases

import (
	"context"
	dtos "go-api/domain/dtos/album"
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

func (uc *GetAverageAlbumPricesUsecase) Execute(ctx context.Context) dtos.GetAverageAlbunsPriceResponseDto {
	albumsPtr := uc.Repo.GetAllWithoutPagination(ctx)
	if len(*albumsPtr) == 0 || albumsPtr == nil {
		zero := float64(0)
		return dtos.GetAverageAlbunsPriceResponseDto{AveragePrice: &zero}
	}

	albums := *albumsPtr
	var totalSum int64 = 0

	for _, a := range albums {
		totalSum += int64(a.Price)
	}

	average := float64(totalSum) / float64(len(albums))
	response := dtos.GetAverageAlbunsPriceResponseDto{
		AveragePrice: &average,
	}
	return response
}
