package database

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/Artragnus/go-personal-finance-app/internal/entity"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{DB: db}
}

func (u *User) Create(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *User) GetByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := u.DB.Where("email = ?", email).First(&user).Error(); err != nil {
		return &entity.User{}, err
	}

	fmt.Println(user)

	return &entity.User{}, nil
}
