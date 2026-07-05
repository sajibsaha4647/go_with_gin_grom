package user

import (
	"ecommerce/config"
	"ecommerce/rest/middleware"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) RegisterUserRoutes(router *gin.Engine, cfg *config.Config) {

	api := router.Group("/api")
	userRouter := api.Group("/user")

	// Public
	userRouter.POST("/login", h.Login)
	userRouter.POST("/register", h.register)

	userRouter.Use(
		middleware.AuthenticationMiddleware(cfg.JWTSecret),
	)


	userRouter.GET("/", h.userlist)
	userRouter.GET("/:id", h.singleUser)
	userRouter.PUT("/:id", h.updateUser)
	userRouter.DELETE("/:id", h.deleteUser)
	userRouter.GET("/count", h.rowCount)

}
