package dto

type CategoryExpense struct {
	ID          int64  `gorm:"autoIncrement" json:"id"`
	Description string `                     json:"Description"`
}

type CategoryIncome struct {
	ID          int64  `gorm:"autoIncrement" json:"id"`
	Description string `                     json:"Description"`
}
