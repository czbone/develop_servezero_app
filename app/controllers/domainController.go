package controllers

import (
	"net/http"
	"strings"
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
	// 入力値取得
	name := strings.TrimSpace(c.PostForm("name")) // ドメイン名

	// ドメイン取得
	domainDb := &db.DomainDb{}
	rows := domainDb.GetDomainList()

	// ドメイン存在確認
	row := domainDb.GetDomainByName(name)
	if row != nil {
		c.HTML(http.StatusOK, "domain.tmpl.html", pongo2.Context{
			"title":      config.GetEnv().AppName, // ナビゲーションメニュータイトル
			"page_title": "ドメイン一覧",
			"domainList": rows,
			"error":      "登録済みです",
		})
		return
	}

	// ドメイン一覧表示
	c.HTML(http.StatusOK, "domain.tmpl.html", pongo2.Context{
		"title":      config.GetEnv().AppName, // ナビゲーションメニュータイトル
		"page_title": "ドメイン一覧",
		"domainList": rows,
	})
}
