package dtos

type UpdateAlbumDto struct {
	Title  *string  `json:"title" validate:"omitempty,min=3"`
	Artist *string  `json:"artist" validate:"omitempty"`
	Price  *float64 `json:"price" validate:"omitempty,gte=0"`
}
