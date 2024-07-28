package router

import (
	"net/http"
	"xpm-auth/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter(authController *controller.AuthController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api")
	baseRouter.POST("/register", authController.Register)
	baseRouter.GET("/login", authController.Login)

	return router
}
