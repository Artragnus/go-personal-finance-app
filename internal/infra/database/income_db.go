package database

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/Artragnus/go-personal-finance-app/internal/dto"
	"github.com/Artragnus/go-personal-finance-app/internal/entity"
	"github.com/Artragnus/go-personal-finance-app/internal/token"
)

type Income struct {
	DB *gorm.DB
}

func NewIncome(db *gorm.DB) *Income {
	return &Income{
		DB: db,
	}
}

func (i *Income) Create(
	income *entity.Income,
) (dto.IncomeCreateResponse, error) {
	var res dto.IncomeCreateResponse
	if err := i.DB.Create(&income).Clauses(clause.Returning{}).Scan(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (i *Income) Get(payload token.Payload) ([]dto.IncomeResponse, error) {
	var res []dto.IncomeResponse
	if err := i.DB.Where("user_id = ?", payload.ID).Find(&[]entity.Income{}).Scan(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}
