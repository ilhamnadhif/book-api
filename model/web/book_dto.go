package web

import "time"

type (
	BookResponse struct {
		ID          int       `json:"id"`
		UserID      int       `json:"user_id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Price       int       `json:"price"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	BookCreateReq struct {
		UserID      int    `json:"user_id" validate:"required"`
		Name        string `json:"name" validate:"required"`
		Description string `json:"description" validate:"required"`
		Price       int    `json:"price" validate:"required"`
	}

	BookUpdateReq struct {
		ID          int    `json:"id" validate:"required"`
		Name        string `json:"name" validate:"required"`
		Description string `json:"description" validate:"required"`
		Price       int    `json:"price" validate:"required"`
	}
)
