package product

type ProductRepository struct {
	 db *gorm.DB
}


func NewProductRepository(db *gorm.DB) *ProductPort { //in this way interface here implemented by struct ProductRepository and return type is ProductPort interface
	return &ProductRepository{db: db}
}



func (r *ProductRepository) createProduct(product domain.Product) error {
	return r.db.Create(&product).Error
}