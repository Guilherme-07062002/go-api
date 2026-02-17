package dtos

type UpdateAlbumDto struct {
	Title  *string  `json:"title" validate:"omitempty,min=3" example:"Blue Train"`
	Artist *string  `json:"artist" validate:"omitempty" example:"John Coltran"`
	Price  *float64 `json:"price" validate:"omitempty,gte=0" example:"56.99"`
}
