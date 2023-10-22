package product

import (
	"context"
	"log"
)

type Repository interface {
	Create(ctx context.Context, req Product) (err error)
	GetAll(ctx context.Context) ([]Product, error)
	GetByID(ctx context.Context, ID int) (Product, error)
}

type Service struct {
	repo PostgreGormRepository
}

func NewService(repo PostgreGormRepository) Service {
	return Service{repo: repo}
}

func (s Service) CreateProduct(ctx context.Context, req Product) (err error) {
	if err = req.Validate(); err != nil {
		log.Println("error when try to validate request with error :", err.Error())
		return
	}

	if err = s.repo.Create(ctx, req); err != nil {
		log.Println("error when try to Create to database with error :", err.Error())
		return

	}

	return
}

func (s Service) GetProducts(ctx context.Context) (products []Product, err error) {
	products, err = s.repo.GetAll(ctx)
	if err != nil {
		log.Println("error when get all products with error :", err.Error())
	}
	return
}

func (s Service) GetProduct(ctx context.Context, ID int) (product Product, err error) {
	product, err = s.repo.GetByID(ctx, ID)
	if err != nil {
		log.Println("error when get product with error :", err.Error())
	}
	return
}
