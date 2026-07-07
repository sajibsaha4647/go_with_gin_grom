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
	if !utils.BindAndValidateFromData(c, &req) {
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

	idStr := c.Param("id")

	prod, err := p.repo.SingleProduct(idStr)
	if err != nil {
		utils.SendError(c, 404, "User not found")
		return
	}

	utils.SendSuccess(c, 200, "User found", prod)
}

func (p *ProductHandler) GetAllProducts(c *gin.Context) {

	userList, err := p.repo.Productlist()
	if err != nil {
		utils.SendError(c, 400, "No product found")
		return
	}

	utils.SendSuccess(c, 200, "product list found", userList)

}

func (p *ProductHandler) UpdateProduct(c *gin.Context) {

	id := c.Param("id")

	var req productRequest.ProductCreateRequest

	if !utils.BindAndValidate(c, &req) {
		return
	}

	// Get existing product
	product, err := p.repo.SingleProduct(id)
	if err != nil {
		utils.SendError(c, http.StatusNotFound, "Product not found")
		return
	}

	// Update fields
	product.Title = req.Title
	product.Description = req.Description
	product.Price = req.Price

	// Optional image upload
	file, err := c.FormFile("image")
	if err == nil {

		imageName, err := utils.UploadImageFile(file, "products")
		if err != nil {
			utils.SendError(c, http.StatusBadRequest, err.Error())
			return
		}

		product.Image = imageName
		product.ImageURL = "/uploads/products/" + imageName
	}

	updatedProduct, err := p.repo.UpdateProduct(id, product)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendSuccess(c, 200, "Product updated successfully", updatedProduct)
}

func (p *ProductHandler) DeleteProduct(c *gin.Context) {

	id := c.Param("id")

	err := p.repo.DeleteProduct(id)
	if err != nil {
		utils.SendError(c, 400, "Delete failed")
		return
	}

	utils.SendSuccess(c, 200, "Successfully Deleted", nil)

}

func (p *ProductHandler) RowCount(c *gin.Context) {

	count, err := p.repo.RowCount()
	if err != nil {
		utils.SendError(c, 400, "count not found")
		return
	}

	utils.SendSuccess(c, 200, "Row count found", count)

}
