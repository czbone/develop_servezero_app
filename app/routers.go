package main

import (
	"net/http"
	"web/config"
	"web/modules/filters/auth"
	routeRegister "web/routes"

	//"github.com/gin-contrib/pprof"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	"github.com/stnc/pongo2gin"
)

func initRouter() *gin.Engine {
	// システムのデフォルトのミドルウェア(ログ機能、リカバリー機能)は組み込まない
	router := gin.New()

	// テンプレートエンジン設定
	router.HTMLRender = pongo2gin.TemplatePath("templates")

	// Javascriptファイル、CSSファイル、画像ファイルを公開
	router.Static("/assets", "public/assets")

	// 処理時間計測用
	// if config.GetEnv().Debug {
	// 	pprof.Register(router)
	// }

	// ミドルウェアの設定
	router.Use(gin.Logger())
	//router.Use(gin.Recovery())
	router.Use(handleErrors()) // リカバリー機能
	router.Use(auth.RegisterGlobalAuthDriver("file" /*ドライバータイプ*/, "web_auth" /*認証タイプ*/))

	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.tmpl.html",
			pongo2.Context{
				"language": config.GetEnv().DefaultLanguage,
				"title":    "404 - ページが見つかりません",
				"message":  "ページが見つかりません",
			},
		)
	})

	router.NoMethod(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.tmpl.html",
			pongo2.Context{
				"language": config.GetEnv().DefaultLanguage,
				"title":    "404 - ページが見つかりません",
				"message":  "ページが見つかりません",
			},
		)
	})

	// ルーティング設定
	routeRegister.RegisterApiRouter(router)
	routeRegister.RegisterPageRouter(router)

	return router
}
