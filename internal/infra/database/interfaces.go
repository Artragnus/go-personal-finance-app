package database

import "github.com/Artragnus/go-personal-finance-app/internal/entity"

type UserInterface interface { 
	Create(user *entity.User) error
	
}