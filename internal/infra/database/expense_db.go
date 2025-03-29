package database

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/Artragnus/go-personal-finance-app/internal/dto"
	"github.com/Artragnus/go-personal-finance-app/internal/entity"
)

type Expense struct {
	DB *gorm.DB
}

func NewExpense(db *gorm.DB) *Expense {
	return &Expense{
		DB: db,
	}
}

func (e *Expense) Create(expense *entity.Expense) (dto.ExpenseCreateResponse, error) {
	var res dto.ExpenseCreateResponse
	if err := e.DB.Create(&expense).Clauses(clause.Returning{}).Scan(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}
