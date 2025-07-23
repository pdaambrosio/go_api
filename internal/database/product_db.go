package database

import (
	"api/internal/entity"

	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{DB: db}
}

func (p *Product) Create(product *entity.Product) error {
	if err := p.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func (p *Product) FindByID(id string) (*entity.Product, error) {
	var product entity.Product
	if err := p.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *Product) FindAll() ([]*entity.Product, error) {
	var products []*entity.Product
	if err := p.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p *Product) Update(product *entity.Product) error {
	if err := p.DB.Save(product).Error; err != nil {
		return err
	}
	return nil
}

func (p *Product) Delete(id string) error {
	if err := p.DB.Where("id = ?", id).Delete(&entity.Product{}).Error; err != nil {
		return err
	}
	return nil
}
