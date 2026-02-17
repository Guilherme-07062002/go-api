package usecases_test

import (
	"context"
	dtos "go-api/domain/dtos/album"
	"go-api/domain/entities"
	"go-api/usecases"
	"go-api/usecases/tests/mocks"
	"math"
	"testing"
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

			// Comparar os valores dos ponteiros, não os endereços
			if result.AveragePrice == nil && tt.expected.AveragePrice == nil {
				return // ambos são nil, teste passou
			}
			if result.AveragePrice == nil || tt.expected.AveragePrice == nil {
				t.Errorf("Executou %s: esperado %v, mas obteve %v", tt.name, tt.expected, result)
				return
			}
			// Comparar os valores com pequena tolerância para ponto flutuante
			if math.Abs(*result.AveragePrice-*tt.expected.AveragePrice) > 0.0001 {
				t.Errorf("Executou %s: esperado %v, mas obteve %v", tt.name, *tt.expected.AveragePrice, *result.AveragePrice)
			}
		})
	}
}
