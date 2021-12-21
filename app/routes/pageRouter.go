package routes

import (
	"web/controllers"
	"web/modules/filters/auth"

	"github.com/gin-gonic/gin"
)

func RegisterPageRouter(router *gin.Engine) {
	controller := &controllers.PageController{}
	router.POST("/login", controller.Login)
	router.GET("/logout", controller.Logout)

	router.Use(auth.Middleware(auth.FileAuthDriverKey))
	{
		router.GET("/", controller.Index)
	}
}
