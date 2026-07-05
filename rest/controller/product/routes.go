package product

import (
	"ecommerce/config"
	"ecommerce/rest/middleware"

	"github.com/gin-gonic/gin"
)

func (h *ProductHandler) RegisterProductRoutes(router *gin.Engine, cfg *config.Config) {

	// Protected routes
	api := router.Group("/api")
	product := api.Group("/product")

	product.Use(
		middleware.AuthenticationMiddleware(cfg.JWTSecret),
	)

	product.GET("/", h.GetAllProducts)
	product.GET("/:id", h.GetProductByID)
	product.POST("/", h.CreateProduct)
	product.PUT("/:id", h.UpdateProduct)
	product.DELETE("/:id", h.DeleteProduct)
	product.GET("/count", h.RowCount)

}
