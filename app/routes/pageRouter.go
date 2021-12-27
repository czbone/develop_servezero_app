package routes

import (
	"web/controllers"
	"web/modules/filters/auth"

	"github.com/gin-gonic/gin"
)

func RegisterPageRouter(router *gin.Engine) {
	// コントローラ作成
	loginController := &controllers.LoginController{}
	domainController := &controllers.DomainController{}

	router.POST("/login", loginController.Login)
	router.GET("/logout", loginController.Logout)

	// その他のページは認証後アクセス可能
	router.Use(auth.Middleware(auth.FileAuthDriverKey))
	{
		router.GET("/", domainController.Index)
	}
}
