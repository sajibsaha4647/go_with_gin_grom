package rest

import (
	"ecommerce/config"
	"ecommerce/rest/controller/product"
	"ecommerce/rest/controller/user"
	"ecommerce/rest/middleware"
	"fmt"
	"net/http"

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

	// setMiddleware.Use(
	// 	middleware.LoggerMiddleware(),
	// 	middleware.CORSMiddleware(),
	// )

	server.Use(setMiddleware.Middlewares()...)

	s.productHandler.RegisterProductRoutes(server, cfg)
	s.userHandler.RegisterUserRoutes(server, cfg)

	fmt.Println("server running on " + cfg.Port)

	if err := http.ListenAndServe(":"+cfg.Port, server); err != nil {
		fmt.Println(err, "error from server")
	}

}
