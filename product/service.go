package product

import "ecommerce/domain"

type ProductServe interface {
	Store(product *domain.Product) error
	SingleProduct(id string) (domain.Product, error)
	Productlist() ([]domain.Product, error)
	UpdateProduct(id string, product domain.Product) error
	DeleteProduct(id string) error
	RowCount() (int64, error)
}

type productService struct {
	repository ProductPort
}

func NewProductService(repository ProductPort) ProductServe {
	return &productService{repository: repository}
}

// deleteProduct implements [ProductServe].
func (p *productService) DeleteProduct(id string) error {

	return p.repository.deleteProduct(id)
}
	

// productlist implements [ProductServe].
func (p *productService) Productlist() ([]domain.Product, error) {
	return p.repository.getAllProducts()
}

// rowCount implements [ProductServe].
func (p *productService) RowCount() (int64, error) {
	return p.repository.rowCount()

}

// singleProduct implements [ProductServe].
func (p *productService) SingleProduct(id string) (domain.Product, error) {
	return p.repository.getProductByID(id)
}

// store implements [ProductServe].
func (p *productService) Store(product *domain.Product) error {
	return p.repository.createProduct(product)
}


// updateProduct implements [ProductServe].
func (p *productService) UpdateProduct(id string, product domain.Product) error {
	return p.repository.updateProduct(id, product)
}


