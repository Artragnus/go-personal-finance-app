package database

import (
	"github.com/Artragnus/go-personal-finance-app/internal/dto"
	"github.com/Artragnus/go-personal-finance-app/internal/entity"
)

type UserInterface interface {
	Create(user *entity.User) error
	GetByEmail(email string) (entity.User, error)
}

type ExpenseInterface interface {
	Create(expense *entity.Expense) (dto.ExpenseCreateResponse, error)
}

type IncomeInterface interface {
	Create(income *entity.Income) (dto.IncomeCreateResponse, error)
}
