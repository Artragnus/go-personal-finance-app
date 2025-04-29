package dto

import (
	"time"

	"github.com/Artragnus/go-personal-finance-app/pkg/entity"
)

type ExpenseCreateRequest struct {
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
	CategoryID  int64  `json:"category_id"`
}

type ExpenseCreateResponse struct {
	ID          entity.ID `json:"id"`
	UserID      entity.ID `json:"user_id"`
	Amount      int64     `json:"amount"`
	Description string    `json:"description"`
	CategoryID  int64     `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type ExpenseResponse struct {
	ID          entity.ID `json:"id"`
	Amount      int64     `json:"amount"`
	Description string    `json:"description"`
	CategoryID  int64     `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
}
