package dtos

type CreateAlbumDto struct {
	Title  string  `json:"title" binding:"required" validate:"required,min=3"`
	Artist string  `json:"artist" binding:"required" validate:"required"`
	Price  float64 `json:"price" binding:"required" validate:"gte=0"`
}
