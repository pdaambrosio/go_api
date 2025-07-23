package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Connor", "john@example.com", "password123")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "John Connor", user.Username)
	assert.Equal(t, "john@example.com", user.Email)
	assert.NotEqual(t, "password123", user.Password)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, _ := NewUser("John Connor", "john@example.com", "password123")
	assert.True(t, user.ValidatePassword("password123"))
	assert.False(t, user.ValidatePassword("wrongpassword"))
	assert.NotEqual(t, "password123", user.Password) // Ensure password is hashed
}
