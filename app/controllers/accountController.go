package controllers

import (
	"net/http"
	"web/config"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

// ######################################################################
// ページコントローラ
// ・GETまたはPOSTでHTTPリクエストを受信し、WEB画面を作成して、HTMLデータを返す
// ######################################################################
type AccountController struct{}

func (pc *AccountController) Index(c *gin.Context) {
	// ドメイン一覧表示
	c.HTML(http.StatusOK, "account.tmpl.html", pongo2.Context{
		"app_name":   config.GetEnv().AppName, // ナビゲーションメニュータイトル
		"page_title": "アカウント",
	})
}
