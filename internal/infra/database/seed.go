package database

import (
	"log"

	"gorm.io/gorm"

	"github.com/Artragnus/go-personal-finance-app/internal/entity"
)

func Seed(db *gorm.DB) {
	categoryExpenses := []entity.CategoryExpense{
		{ID: 1, Description: "Casa"},
		{ID: 2, Description: "Educação"},
		{ID: 3, Description: "Eletrônicos"},
		{ID: 4, Description: "Lazer"},
		{ID: 5, Description: "Outros"},
		{ID: 6, Description: "Restaurante"},
		{ID: 7, Description: "Saúde"},
		{ID: 8, Description: "Serviço"},
		{ID: 9, Description: "Supermercado"},
		{ID: 10, Description: "Transporte"},
		{ID: 11, Description: "Vestuário"},
		{ID: 12, Description: "Viagem"},
	}
	var categoryExpensesCount int64

	if err := db.Model(entity.CategoryExpense{}).Count(&categoryExpensesCount).Error; err != nil {
		log.Println("error to get category expenses")
	}

	switch categoryExpensesCount {
	default:
		for i := categoryExpensesCount; int(i) < len(categoryExpenses); i++ {
			if err := db.Create(&categoryExpenses[i]).Error; err != nil {
				log.Println("error to seed category expenses")
			}
		}
	case 0:
		if err := db.Create(&categoryExpenses).Error; err != nil {
			log.Println("error to seed category expenses")
		}
	case int64(len(categoryExpenses)):
		return
	}
}
