package entity

import (
	"errors"
	"time"

	"github.com/Artragnus/go-personal-finance-app/pkg/entity"
)

type Expense struct {
	ID          entity.ID `json:"id"`
	UserID      entity.ID `json:"user_id"`
	Amount      int64     `json:"amount"`
	Description string    `json:"description"`
	CategoryID  int64     `json:"category_id"`
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
	expense := &Expense{
		ID:          entity.NewID(),
		UserID:      userID,
		Amount:      amount,
		Description: description,
		CategoryID:  categoryID,
		Status:      true,
	}

	if err := expense.Validate(); err != nil {
		return expense, err
	}

	return expense, nil
}

func (e *Expense) Validate() error {
	if e.Amount < 1 {
		return errors.New("the amount should be more than 0")
	}
	return nil
}
