package product

type ProductFormatter struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
	Stock    int    `json:"stock"`
}

func FormatProduct(product Product) ProductFormatter {
	productFormatter := ProductFormatter{}

	productFormatter.Id = product.Id
	productFormatter.Name = product.Name
	productFormatter.Category = product.Category
	productFormatter.Price = product.Price
	productFormatter.Stock = product.Stock

	return productFormatter
}

func FormatProducts(products []Product) []ProductFormatter {
	productsFormatter := []ProductFormatter{}

	for _, product := range products {
		productFormatter := FormatProduct(product)
		productsFormatter = append(productsFormatter, productFormatter)
	}

	return productsFormatter
}
