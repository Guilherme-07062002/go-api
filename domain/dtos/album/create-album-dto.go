package dtos

type CreateAlbumDto struct {
	Title  string  `json:"title" validate:"required,min=3" example:"Blue Train"`
	Artist string  `json:"artist" validate:"required" example:"John Coltran"`
	Price  float64 `json:"price" validate:"required,gt=0" example:"56.99"`
}
