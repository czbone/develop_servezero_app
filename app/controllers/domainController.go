package controllers

import (
	"net/http"
	"web/config"
	"web/db"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

// ######################################################################
// ページコントローラ
// ・GETまたはPOSTでHTTPリクエストを受信し、WEB画面を作成して、HTMLデータを返す
// ######################################################################
type DomainController struct{}

func (pc *DomainController) Index(c *gin.Context) {
	// ドメイン取得
	domainDb := &db.DomainDb{}
	rows := domainDb.GetDomainList()

	// ドメイン一覧表示
	c.HTML(http.StatusOK, "domain.tmpl.html", pongo2.Context{
		"title":      config.GetEnv().AppName, // ナビゲーションメニュータイトル
		"page_title": "ドメイン一覧",
		"domainList": rows,
	})
}
