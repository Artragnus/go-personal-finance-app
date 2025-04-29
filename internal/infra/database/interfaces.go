package database

import (
	"github.com/Artragnus/go-personal-finance-app/internal/dto"
	"github.com/Artragnus/go-personal-finance-app/internal/entity"
	"github.com/Artragnus/go-personal-finance-app/internal/token"
)

type UserInterface interface {
	Create(user *entity.User) error
	GetByEmail(email string) (entity.User, error)
}

type ExpenseInterface interface {
	Create(expense *entity.Expense) (dto.ExpenseCreateResponse, error)
	Get(payload token.Payload) ([]dto.ExpenseResponse, error)
}

type IncomeInterface interface {
	Create(income *entity.Income) (dto.IncomeCreateResponse, error)
	Get(payload token.Payload) ([]dto.IncomeResponse, error)
}

type CategoryInterface interface {
	GetIncomeCategories() ([]dto.CategoryIncome, error)
	GetExpenseCategories() ([]dto.CategoryExpense, error)
}
