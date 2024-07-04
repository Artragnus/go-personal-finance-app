package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "password", "teste@teste.com")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "teste@teste.com", user.Email)
} 

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe", "password", "teste@teste.com")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("password"))
	assert.False(t, user.ValidatePassword("password123"))
	assert.NotEqual(t, "password", user.Password)
	
}