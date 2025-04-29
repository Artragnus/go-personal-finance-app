package dto

import (
	"time"

	"github.com/Artragnus/go-personal-finance-app/pkg/entity"
)

type IncomeCreateRequest struct {
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
	CategoryID  int64  `json:"category_id"`
}

type IncomeCreateResponse struct {
	ID          entity.ID `json:"id"`
	UserID      entity.ID `json:"user_id"`
	Amount      int64     `json:"amount"`
	Description string    `json:"description"`
	CategoryID  int64     `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
}
type IncomeResponse struct {
	ID          entity.ID `json:"id"`
	Amount      int64     `json:"amount"`
	Description string    `json:"description"`
	CategoryID  int64     `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
}
