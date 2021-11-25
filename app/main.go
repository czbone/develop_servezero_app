package main

import (
	//"github.com/gin-gonic/gin"

	"runtime"
	//"web/config"
	"web/modules/server"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// ### Ginのモード変更 ###
	/*if config.GetEnv().Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode) // リリースモードに設定([Gin-debug]出力オフ)
	}*/

	// ルーティング設定
	router := initRouter()

	// Webサーバ起動
	server.Run(router)
}
