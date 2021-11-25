package main

import (
	"net/http"
	"path/filepath"

	//"web/filters"
	//"web/filters/auth"
	"web/config"
	routeRegister "web/routes"

	//"github.com/gin-contrib/pprof"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.New()

	// テンプレートファイルを読み込む
	//router.HTMLRender = loadTemplates("templates")
	//router.LoadHTMLGlob(config.GetEnv().TemplatePath + "/*.tmpl") // テンプレートディレクトリ
	//router.LoadHTMLGlob(config.GetEnv().TemplatePath + "/*/*.tmpl") // テンプレートディレクトリ
	router.LoadHTMLGlob("templates/*/*")

	// Javascriptファイル、CSSファイル、画像ファイルを公開
	router.Static("/assets", "public/assets")

	// 処理時間計測用
	// if config.GetEnv().Debug {
	// 	pprof.Register(router)
	// }

	// ミドルウェアの設定
	router.Use(gin.Logger())

	router.Use(handleErrors())
	//router.Use(filters.RegisterSession())
	//router.Use(filters.RegisterCache())

	//router.Use(auth.RegisterGlobalAuthDriver("cookie", "web_auth"))
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
func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.tmpl")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/includes/*.tmpl")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
