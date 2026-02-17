package usecases_test

import (
	"context"
	dtos "go-api/domain/dtos/album"
	"go-api/domain/entities"
	"go-api/usecases"
	"go-api/usecases/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAverageAlbumPricesUsecase_Execute(t *testing.T) {
	avg20 := 20.0
	avg12_5 := 12.5
	avg0 := 0.0

	tests := []struct {
		name     string
		mockData []entities.Album
		expected dtos.GetAverageAlbunsPriceResponseDto
	}{
		{
			name: "Média simples de três álbuns",
			mockData: []entities.Album{
				{Price: 10}, {Price: 20}, {Price: 30},
			},
			expected: dtos.GetAverageAlbunsPriceResponseDto{AveragePrice: &avg20},
		},
		{
			name: "Média com valores decimais",
			mockData: []entities.Album{
				{Price: 10}, {Price: 15},
			},
			expected: dtos.GetAverageAlbunsPriceResponseDto{AveragePrice: &avg12_5},
		},
		{
			name:     "Lista vazia deve retornar zero",
			mockData: []entities.Album{},
			expected: dtos.GetAverageAlbunsPriceResponseDto{AveragePrice: &avg0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocks.MockAlbumRepository{
				MockGetAllWithoutPagination: func(ctx context.Context) *[]entities.Album {
					return &tt.mockData
				},
			}

			uc := usecases.NewGetAverageAlbumPricesUsecase(mockRepo)

			result := uc.Execute(context.Background())

			assert.Equal(t, tt.expected, result, "O DTO de retorno deve ser idêntico ao esperado")
		})
	}
}
