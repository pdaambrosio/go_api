package database

import (
	"api/internal/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	db.AutoMigrate(&entity.Product{})

	productDB := NewProduct(db)

	product, err := entity.NewProduct("Test Product", 10.0, "This is a test product")
	assert.NoError(t, err)
	err = productDB.Create(product)
	assert.NoError(t, err)

	fetchedProduct, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.Name, fetchedProduct.Name)
	assert.Equal(t, product.Price, fetchedProduct.Price)
	assert.Equal(t, product.Description, fetchedProduct.Description)
	// Compare timestamps by truncating to seconds to avoid timezone issues
	assert.True(t, product.CreatedAt.Truncate(time.Second).Equal(fetchedProduct.CreatedAt.Truncate(time.Second)))
	assert.True(t, product.UpdatedAt.Truncate(time.Second).Equal(fetchedProduct.UpdatedAt.Truncate(time.Second)))
	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)
	_, err = productDB.FindByID(product.ID.String())
	assert.Error(t, err, "expected error when fetching deleted product")
}
