package database

import (
	"github.com/Artragnus/go-personal-finance-app/internal/dto"
	"github.com/Artragnus/go-personal-finance-app/internal/entity"
	"gorm.io/gorm"
)

type Category struct {
	DB *gorm.DB
}

func NewCategory(db *gorm.DB) *Category {
	return &Category{
		DB: db,
	}
}

func (c *Category) GetIncomeCategories() ([]dto.CategoryIncome, error) {
	var categories []dto.CategoryIncome

	if err := c.DB.Find(&entity.CategoryIncome{}).Scan(&categories).Error; err != nil {
		return categories, err
	}

	return categories, nil
}

func (c *Category) GetExpenseCategories() ([]dto.CategoryExpense, error) {
	var categories []dto.CategoryExpense

	if err := c.DB.Find(&entity.CategoryExpense{}).Scan(&categories).Error; err != nil {
		return categories, err
	}

	return categories, nil
}
