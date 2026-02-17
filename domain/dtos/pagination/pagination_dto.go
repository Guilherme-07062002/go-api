package dtos

type PaginationDto struct {
	Page  int `form:"page,default=1" json:"page" validate:"min=1"`
	Limit int `form:"limit,default=10" json:"limit" validate:"min=1,max=100"`
}

type PaginatedResponse[T any] struct {
	Data       []T   `json:"data"`
	Total      int64 `json:"total"`
	Page       int64 `json:"page"`
	Limit      int64 `json:"limit"`
	TotalPages int64 `json:"total_pages"`
}
