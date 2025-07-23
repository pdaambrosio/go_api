package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduct_Validate(t *testing.T) {
	prod, err := NewProduct("Test Product", 10.0, "A test product description")
	assert.Nil(t, err)
	assert.NotNil(t, prod)
	assert.Equal(t, "Test Product", prod.Name)
	assert.Equal(t, 10.0, prod.Price)
	assert.NotEmpty(t, prod.CreatedAt)
	assert.NotEmpty(t, prod.UpdatedAt)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	prod, err := NewProduct("", 10.0, "A test product description")
	assert.NotNil(t, err)
	assert.Nil(t, prod)
	assert.Equal(t, ErrNameRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	prod, err := NewProduct("Test Product", 0.0, "A test product description")
	assert.NotNil(t, err)
	assert.Nil(t, prod)
	assert.Equal(t, ErrPriceMustBePositive, err)
}

func TestProductWhenPriceIsInvalide(t *testing.T) {
	prod, err := NewProduct("Test Product", -5.0, "A test product description")
	assert.NotNil(t, err)
	assert.Nil(t, prod)
	assert.Equal(t, ErrPriceMustBePositive, err)
}

func TestProductValidate(t *testing.T) {
	prod, err := NewProduct("Test Product", 10.0, "A test product description")
	assert.Nil(t, err)
	assert.NotNil(t, prod)
	assert.Equal(t, "Test Product", prod.Name)
	assert.Equal(t, 10.0, prod.Price)
	assert.NotEmpty(t, prod.CreatedAt)
	assert.NotEmpty(t, prod.UpdatedAt)
}
