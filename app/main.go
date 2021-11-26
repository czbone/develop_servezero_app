// コマンド起動で最初に呼ばれるmain処理です。
// リリース時は環境変数GIN_MODEにreleaseを設定してリリースモードでコマンド実行します。リリースモードではデバッグ用のコンソール出力が抑止されます。
// export GIN_MODE=release
package main

import (
	"runtime"
	"web/modules/server"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// ルーティング設定
	router := initRouter()

	// Webサーバ起動
	server.Run(router)
}
