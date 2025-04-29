package database

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/Artragnus/go-personal-finance-app/internal/dto"
	"github.com/Artragnus/go-personal-finance-app/internal/entity"
	"github.com/Artragnus/go-personal-finance-app/internal/token"
)

type Expense struct {
	DB *gorm.DB
}

func NewExpense(db *gorm.DB) *Expense {
	return &Expense{
		DB: db,
	}
}

func (e *Expense) Create(
	expense *entity.Expense,
) (dto.ExpenseCreateResponse, error) {
	var res dto.ExpenseCreateResponse
	if err := e.DB.Create(&expense).Clauses(clause.Returning{}).Scan(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (e *Expense) Get(payload token.Payload) ([]dto.ExpenseResponse, error) {
	var res []dto.ExpenseResponse
	if err := e.DB.Where("user_id = ?", payload.ID).Find(&[]entity.Expense{}).Scan(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}
