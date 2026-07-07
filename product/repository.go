package product

import (
	"ecommerce/domain"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductPort { //in this way interface here implemented by struct ProductRepository and return type is ProductPort interface
	return &ProductRepository{db: db}
}

func (r *ProductRepository) createProduct(product *domain.Product) error {
	return r.db.Create(&product).Error
}

func (r *ProductRepository) getProductByID(id string) (domain.Product, error) {
	var product domain.Product
	err := r.db.First(&product, id).Error
	return product, err
}

func (r *ProductRepository) getAllProducts() ([]domain.Product, error) {
	var products []domain.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *ProductRepository) updateProduct(id string, product domain.Product) error {
	return r.db.Model(&product).Where("id = ?", id).Updates(product).Error
}

func (r *ProductRepository) deleteProduct(id string) error {
	var product domain.Product
	return r.db.Where("id = ?", id).Delete(&product).Error
}

func (r *ProductRepository) rowCount() (int64, error) {
	var count int64
	err := r.db.Model(&domain.Product{}).Count(&count).Error
	return count, err
}