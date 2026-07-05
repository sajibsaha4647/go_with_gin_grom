package rest

import (
	"ecommerce/config"
	"ecommerce/rest/controller/product"
	"ecommerce/rest/controller/user"
	"ecommerce/rest/middleware"

	"github.com/gin-gonic/gin"
)

type Server struct {
	productHandler *product.ProductHandler
	userHandler    *user.UserHandler
}

func NewServer(productHandler *product.ProductHandler, userHandler *user.UserHandler) *Server {
	return &Server{
		productHandler: productHandler,
		userHandler:    userHandler,
	}
}

func (s *Server) Start(cfg *config.Config) {

	server := gin.Default()

	setMiddleware := middleware.NewMiddlewareManager()

	setMiddleware.Use(
		middleware.LoggerMiddleware(),
		middleware.CORSMiddleware(),
	)

	server.Use(setMiddleware.Middlewares()...)

}
