package routes

import (
	"web/controllers"
	"web/modules/filters/auth"

	"github.com/gin-gonic/gin"
)

func RegisterApiRouter(router *gin.Engine) {
	apiRouter := router.Group("/api")
	apiRouter.Use(auth.Middleware(auth.FileAuthDriverKey))
	{
		controller := &controllers.ApiController{}
		apiRouter.GET("/index", controller.Index)
	}
}
