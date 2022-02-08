package controllers

import (
	"net/http"
	"web/config"
	"web/modules/filters/auth"
	"web/modules/log"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

// ######################################################################
// ページコントローラ
// ・GETまたはPOSTでHTTPリクエストを受信し、WEB画面を作成して、HTMLデータを返す
// ######################################################################
type AccountController struct{}

func (pc *AccountController) Index(c *gin.Context) {
	// パラメータ初期化
	//var error, success string // メッセージパラメータ
	//userDb := &db.UserDb{}

	// 入力値取得
	//act := strings.TrimSpace(c.PostForm("act")) // 実行操作

	log.Print(auth.GetDefaultUser(c))

	// ドメイン一覧表示
	c.HTML(http.StatusOK, "account.tmpl.html", pongo2.Context{
		"app_name":   config.GetEnv().AppName, // ナビゲーションメニュータイトル
		"page_title": "アカウント",
	})
}
