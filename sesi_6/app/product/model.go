package product

import "errors"

var (
	ErrEmptyName     = errors.New("name tidak boleh kosong")
	ErrEmptyCategory = errors.New("category tidak boleh kosong")
	ErrEmptyPrice    = errors.New("price tidak boleh kosong")
	ErrEmptyStock    = errors.New("stock tidak boleh kosong")
)

type Product struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
	Stock    int    `json:"stock"`
}

func (p Product) Validate() (err error) {
	if p.Name == "" {
		return ErrEmptyName
	}

	if p.Category == "" {
		return ErrEmptyCategory
	}
	if p.Price == 0 {
		return ErrEmptyPrice
	}
	if p.Stock == 0 {
		return ErrEmptyStock
	}
	return
}
