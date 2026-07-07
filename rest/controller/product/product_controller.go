package product

import (
	"ecommerce/domain"
	"ecommerce/dto/productRequest"
	"ecommerce/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *ProductHandler) CreateProduct(c *gin.Context) {

	var req productRequest.ProductCreateRequest
	if !utils.BindAndValidate(c, &req) {
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Failed to retrieve image")
		return
	}

	imageName, err := utils.UploadImageFile(file, "product_images")
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to upload image")
		return
	}

	product := domain.Product{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		Image:       imageName,
		ImageURL:    fmt.Sprintf("/uploads/products/%s", imageName),
	}

	if err := p.repo.Store(&product); err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to create product")
		return
	}

	utils.SendSuccess(c, http.StatusCreated, "Product created successfully", product)

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
