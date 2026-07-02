package product

import "ecommerce/domain"

type ProductServe interface {
	store(product domain.Product) error
	singleProduct(id string) (domain.Product, error)
	productlist() ([]domain.Product, error)
	updateProduct(id string, product domain.Product) error
	deleteProduct(id string) error
	rowCount() (int64, error)
}

type productService struct {
	repository ProductPort
}

func NewProductService(repository ProductPort) ProductServe {
	return &productService{repository: repository}
}

// deleteProduct implements [ProductServe].
func (p *productService) deleteProduct(id string) error {

	return p.repository.deleteProduct(id)
}
	

// productlist implements [ProductServe].
func (p *productService) productlist() ([]domain.Product, error) {
	return p.repository.getAllProducts()
}

// rowCount implements [ProductServe].
func (p *productService) rowCount() (int64, error) {
	return p.repository.rowCount()

}

// singleProduct implements [ProductServe].
func (p *productService) singleProduct(id string) (domain.Product, error) {
	return p.repository.getProductByID(id)
}

// store implements [ProductServe].
func (p *productService) store(product domain.Product) error {
	return p.repository.createProduct(product)
}


// updateProduct implements [ProductServe].
func (p *productService) updateProduct(id string, product domain.Product) error {
	return p.repository.updateProduct(id, product)
}


