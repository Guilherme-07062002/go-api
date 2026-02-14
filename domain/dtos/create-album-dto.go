package dtos

type CreateAlbumDto struct {
	Title  string  `json:"title" validate:"required,min=3"`
	Artist string  `json:"artist" validate:"required"`
	Price  float64 `json:"price" validate:"required,gt=0"`
}
