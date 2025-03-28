package entity

import "github.com/Artragnus/go-personal-finance-app/pkg/entity"

type Expense struct {
	ID          entity.ID `json:"id"`
	UserID      entity.ID `json:"user_id"`
	Amount      int       `json:"amount"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
}

func NewExpense(
	userID entity.ID,
	amount int,
	description string,
	category string,
) (*Expense, error) {
	return &Expense{
		ID:          entity.NewID(),
		UserID:      userID,
		Amount:      amount,
		Description: description,
		Category:    category,
	}, nil
}
