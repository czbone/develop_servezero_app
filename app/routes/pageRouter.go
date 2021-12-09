package routes

import (
	"web/controllers"
	"web/modules/filters/auth"

	"github.com/gin-gonic/gin"
)

func RegisterPageRouter(router *gin.Engine) {
	controller := &controllers.PageController{}
	router.GET("/", controller.Index)
	router.Use(auth.Middleware(auth.CookieAuthDriverKey))
	{
	}
	/*
		api := router.Group("/api")
		api.GET("/index", controllers.IndexApi)
		api.GET("/cookie/set/:userid", controllers.CookieSetExample)

		// cookie auth middleware
		api.Use(auth.Middleware(auth.CookieAuthDriverKey))
		{
			api.GET("/orm", controllers.OrmExample)
			api.GET("/store", controllers.StoreExample)
			api.GET("/db", controllers.DBExample)
			api.GET("/cookie/get", controllers.CookieGetExample)
		}
	*/
}
