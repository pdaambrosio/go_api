package database

import (
	"api/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUser_Create(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("John Connor", "john@example.com", "password123")
	userDB := NewUser(db)
	err = userDB.Create(user)

	assert.NoError(t, err)
	fetchedUser, err := userDB.FindByEmail("john@example.com")
	assert.NoError(t, err)
	assert.Equal(t, user.Email, fetchedUser.Email)
}

func TestUser_FindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{}) // Use an in-memory database for testing
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	db.AutoMigrate(&entity.User{})

	userDB := NewUser(db)

	user, _ := entity.NewUser("John Connor", "john@example.com", "password123")
	err = userDB.Create(user)
	assert.NoError(t, err)

	fetchedUser, err := userDB.FindByEmail("john@example.com")
	assert.NoError(t, err)
	assert.Equal(t, user.Email, fetchedUser.Email)
}
