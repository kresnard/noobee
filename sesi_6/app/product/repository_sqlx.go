package product

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type PostgresSQLXRepository struct {
	db *sqlx.DB
}

func NewPostgresSQLXRepository(db *sqlx.DB) PostgresSQLXRepository {
	return PostgresSQLXRepository{
		db: db,
	}
}

func (p PostgresSQLXRepository) Create(ctx context.Context, model Product) (err error) {
	query := `
		INSERT INTO products (
			name, category, price, stock
		) VALUES (
			:name, :category, :price, :stock
		)
	`

	stmt, err := p.db.PrepareNamed(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(model)

	return
}

func (p PostgresSQLXRepository) GetAll(ctx context.Context) ([]Product, error) {
	var products []Product

	query := `
        SELECT id, name, category, price, stock FROM products
    `

	stmt, err := p.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return products, err
	}
	defer stmt.Close()

	err = p.db.SelectContext(ctx, &products, query)
	if err != nil {
		return products, err
	}
	return products, nil
}

func (p PostgresSQLXRepository) GetByID(ctx context.Context, ID int) (Product, error) {
	var product Product
	query := `
        SELECT id, name, category, price, stock 
		FROM products
		Where id = $1
    `

	stmt, err := p.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return product, err
	}
	defer stmt.Close()

	err = p.db.GetContext(ctx, &product, query, ID)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (p PostgresSQLXRepository) UpdateByID(ctx context.Context, model Product) (err error) {
	query := `
        UPDATE products
		SET name = :name, category = :category, price = :price, stock = :stock
		Where id = :id;
    `

	stmt, err := p.db.PrepareNamed(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(model)
	return
}

func (p PostgresSQLXRepository) DeleteByID(ctx context.Context, ID int) (err error) {
	query := `
        DELETE FROM products
        WHERE id = $1;
    `

	stmt, err := p.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(ID)
	return
}
