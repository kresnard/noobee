package product

import (
	"context"

	"gorm.io/gorm"
)

type PostgreGormRepository struct {
	db *gorm.DB
}

func NewRepositoryGormDB(db *gorm.DB) PostgreGormRepository {
	return PostgreGormRepository{db}
}

func (p PostgreGormRepository) Create(ctx context.Context, model Product) (err error) {
	return p.db.Create(&model).Error
}

func (p PostgreGormRepository) GetAll(ctx context.Context) ([]Product, error) {
	var products []Product

	err := p.db.Find(&products).Error
	if err != nil {
		return products, err
	}

	return products, nil
}

func (p PostgreGormRepository) GetByID(ctx context.Context, ID int) (Product, error) {
	var product Product

	err := p.db.Where("id = ?", ID).Find(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}
