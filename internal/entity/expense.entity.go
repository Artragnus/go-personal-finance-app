package entity

import (
	"time"

	"github.com/Artragnus/go-personal-finance-app/pkg/entity"
)

type Expense struct {
	ID          entity.ID `json:"id"`
	UserID      entity.ID `json:"user_id"`
	Amount      int64     `json:"amount"`
	Description string    `json:"description"`
	CategoryID  int64     `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Status      bool      `json:"status"`
	User        User
	Category    CategoryExpense
}

func NewExpense(
	userID entity.ID,
	amount int64,
	description string,
	categoryID int64,
) (*Expense, error) {
	return &Expense{
		ID:          entity.NewID(),
		UserID:      userID,
		Amount:      amount,
		Description: description,
		CategoryID:  categoryID,
		Status:      true,
	}, nil
}
