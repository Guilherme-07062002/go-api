package mocks

import (
	"go-api/domain/entities"

	uuid "github.com/satori/go.uuid"
)

func GetAlbumsInMemory() []entities.Album {
	return []entities.Album{
		{ID: uuid.NewV4().String(), Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: uuid.NewV4().String(), Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: uuid.NewV4().String(), Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}
}
