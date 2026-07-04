package rest

import (
	"ecommerce/config"
	"ecommerce/rest/controller/product"
	"ecommerce/rest/controller/user"
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

	

}
