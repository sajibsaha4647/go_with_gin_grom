package user

import (
	"ecommerce/config"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) RegisterUserRoutes(router *gin.Engine, cfg *config.Config) {

	userRouter := router.Group("/user")
	userRouter.GET("/", h.userlist)
	userRouter.GET("/:id", h.singleUser)
	userRouter.POST("/", h.store)
	userRouter.PUT("/:id", h.updateUser)
	userRouter.DELETE("/:id", h.deleteUser)
	userRouter.GET("/count", h.rowCount)

}
