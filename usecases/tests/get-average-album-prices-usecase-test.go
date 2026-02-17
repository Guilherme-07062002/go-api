package usecases_test

import (
	"context"
	"go-api/domain/entities"
	"go-api/usecases"
	"testing"
)

func TestGetAverageAlbumPricesUsecase_Execute(t *testing.T) {
	tests := []struct {
		name     string
		mockData []entities.Album
		expected float64
	}{
		{
			name: "Média simples de três álbuns",
			mockData: []entities.Album{
				{Price: 10}, {Price: 20}, {Price: 30},
			},
			expected: 20.0,
		},
		{
			name: "Média com valores decimais",
			mockData: []entities.Album{
				{Price: 10}, {Price: 15},
			},
			expected: 12.5,
		},
		{
			name:     "Lista vazia deve retornar zero",
			mockData: []entities.Album{},
			expected: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockAlbumRepository{
				MockGetAllWithoutPagination: func(ctx context.Context) *[]entities.Album {
					return &tt.mockData
				},
			}

			uc := usecases.NewGetAverageAlbumPricesUsecase(mockRepo)

			result := uc.Execute(context.Background())

			if result != tt.expected {
				t.Errorf("Executou %s: esperado %v, mas obteve %v", tt.name, tt.expected, result)
			}
		})
	}
}
