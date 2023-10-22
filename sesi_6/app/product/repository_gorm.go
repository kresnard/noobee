package product

import (
	"context"

	"gorm.io/gorm"
)

type PostgreGORMRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *PostgreGORMRepository {
	return &PostgreGORMRepository{db}
}

func (p PostgreGORMRepository) Create(ctx context.Context, model Product) (err error) {
	return p.db.Create(&model).Error
}
