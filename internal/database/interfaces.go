package database

import "api/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type ProductInterface interface {
	Create(product *entity.Product) error
	FindByID(id string) (*entity.Product, error)
	FindAll() ([]*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}
