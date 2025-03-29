package entity

import (
	"errors"
	"time"

	"github.com/Artragnus/go-personal-finance-app/pkg/entity"
)

type Income struct {
	ID          entity.ID `json:"id"`
	UserID      entity.ID `json:"user_id"`
	Amount      int64     `json:"amount"`
	Description string    `json:"description"`
	CategoryID  int64     `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Status      bool      `json:"status"`
	User        User
	Category    CategoryIncome
}

func NewIncome(
	userID entity.ID,
	amount int64,
	description string,
	categoryID int64,
) (*Income, error) {
	income := &Income{
		ID:          entity.NewID(),
		UserID:      userID,
		Amount:      amount,
		Description: description,
		CategoryID:  categoryID,
		Status:      true,
	}

	if err := income.Validate(); err != nil {
		return income, err
	}

	return income, nil
}

func (i *Income) Validate() error {
	if i.Amount < 1 {
		return errors.New("the amount should be more than 0")
	}
	return nil
}
