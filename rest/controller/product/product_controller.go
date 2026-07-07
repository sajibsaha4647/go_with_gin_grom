package product

import (
	"ecommerce/dto/productRequest"

	"github.com/gin-gonic/gin"
)

func (p *ProductHandler) CreateProduct(c *gin.Context) {

	var req productRequest.ProductCreateRequest

	if !BindAndValidate(c, &req) {
		return
	}

	



}

func BindAndValidate(c *gin.Context, productCreateRequest *productRequest.ProductCreateRequest) bool {
	panic("unimplemented")
}

func (p *ProductHandler) GetProductByID(c *gin.Context) {
}

func (p *ProductHandler) GetAllProducts(c *gin.Context) {

}

func (p *ProductHandler) UpdateProduct(c *gin.Context) {

}

func (p *ProductHandler) DeleteProduct(c *gin.Context) {
}

func (p *ProductHandler) RowCount(c *gin.Context) {

}
