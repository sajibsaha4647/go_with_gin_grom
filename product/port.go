package product

import "ecommerce/domain"

type ProductPort interface {
	createProduct(product *domain.Product) error
	getProductByID(id string) (domain.Product, error)
	getAllProducts() ([]domain.Product, error)
	updateProduct(id string, product domain.Product) (domain.Product,error)
	deleteProduct(id string) error
	rowCount() (int64, error)
}