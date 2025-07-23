package entity

import (
	"api/pkg/entity"
	"errors"
	"time"
)

var (
	ErrIdRequired          = errors.New("ID is required")
	ErrInvalidId           = errors.New("invalid ID format")
	ErrNameRequired        = errors.New("name is required")
	ErrPriceRequired       = errors.New("price is required")
	ErrPriceMustBePositive = errors.New("price must be greater than zero")
)

type Product struct {
	ID          entity.ID `json:"id"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Description string    `json:"description"`
}

func NewProduct(name string, price float64, description string) (*Product, error) {
	product := &Product{
		ID:          entity.NewID(),
		Name:        name,
		Price:       price,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err := product.Validate()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIdRequired
	}

	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInvalidId
	}

	if p.Name == "" {
		return ErrNameRequired
	}

	if p.Price <= 0 {
		return ErrPriceMustBePositive
	}

	return nil
}
