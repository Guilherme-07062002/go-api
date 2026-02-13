package dtos

type UpdateAlbumDto struct {
	Title  *string  `json:"title,omitempty" validate:"omitempty,required,min=3"`
	Artist *string  `json:"artist,omitempty" validate:"omitempty,required"`
	Price  *float64 `json:"price,omitempty" validate:"omitempty,gte=0"`
}
