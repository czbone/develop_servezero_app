package main

import (
	"net/http"

	"web/config"
	"web/modules/filters"
	"web/modules/filters/auth"
	routeRegister "web/routes"

	//"github.com/gin-contrib/pprof"

	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	//router := gin.New()
	router := gin.Default()

	// テンプレートエンジン設定
	gv := ginview.New(
		goview.Config{
			Root:      "views",
			Extension: ".tmpl",
			Master:    "layouts/master",
		},
	)
	router.HTMLRender = gv

	// Javascriptファイル、CSSファイル、画像ファイルを公開
	router.Static("/assets", "public/assets")

	// 処理時間計測用
	// if config.GetEnv().Debug {
	// 	pprof.Register(router)
	// }

	// ミドルウェアの設定
	router.Use(gin.Logger())

	router.Use(handleErrors())
	router.Use(filters.RegisterSession())

	router.Use(auth.RegisterGlobalAuthDriver("cookie", "web_auth"))
	//router.Use(auth.RegisterGlobalAuthDriver("jwt", "jwt_auth"))

	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.tmpl", gin.H{
			"language": config.GetEnv().DefaultLanguage,
			"title":    "404 - ページが見つかりません",
			"message":  "ページが見つかりません",
		})
	})

	router.NoMethod(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.tmpl", gin.H{
			"language": config.GetEnv().DefaultLanguage,
			"title":    "404 - ページが見つかりません",
			"message":  "ページが見つかりません",
		})
	})

	// ルーティング設定
	routeRegister.RegisterApiRouter(router)
	routeRegister.RegisterPageRouter(router)

	return router
}
