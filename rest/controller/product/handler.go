package product

import "ecommerce/product"

type ProductHandler struct {
	repo product.ProductServe
}


func NewProductHandler(repo product.ProductServe) *ProductHandler {
	return &ProductHandler{repo: repo}
}