package controllers

import (
	"log"
	"net/http"
	"strings"
	"web/config"
	"web/db"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ######################################################################
// ドメイン情報コントローラ
// ・ドメイン情報を管理する
// ######################################################################
type DomainController struct{}

type ValidateNewDomain struct {
	Name string `validate:"fqdn"`
}

func (pc *DomainController) Index(c *gin.Context) {
	// 入力値取得
	act := strings.TrimSpace(c.PostForm("act"))   // 実行操作
	name := strings.TrimSpace(c.PostForm("name")) // ドメイン名

	// ドメイン取得
	domainDb := &db.DomainDb{}
	rows := domainDb.GetDomainList()

	if act == "add" { // ドメイン追加の場合
		// 入力値チェック
		newDomain := &ValidateNewDomain{
			Name: name,
		}
		validate := validator.New()
		err := validate.Struct(newDomain)
		if err != nil {
			log.Print(err)
		}

		// ドメイン存在確認
		row := domainDb.GetDomainByName(name)
		if row == nil {
			// ドメイン追加
			result := domainDb.AddDomain(name, "abc")

			if result {
				c.HTML(http.StatusOK, "domain.tmpl.html", pongo2.Context{
					"title":      config.GetEnv().AppName, // ナビゲーションメニュータイトル
					"page_title": "ドメイン一覧",
					"domainList": rows,
					"success":    "ドメイン登録しました",
				})
			} else {
				c.HTML(http.StatusOK, "domain.tmpl.html", pongo2.Context{
					"title":      config.GetEnv().AppName, // ナビゲーションメニュータイトル
					"page_title": "ドメイン一覧",
					"domainList": rows,
					"error":      "ドメイン登録に失敗しました",
				})
			}
		} else {
			c.HTML(http.StatusOK, "domain.tmpl.html", pongo2.Context{
				"title":      config.GetEnv().AppName, // ナビゲーションメニュータイトル
				"page_title": "ドメイン一覧",
				"domainList": rows,
				"error":      "登録済みのドメインです",
			})
		}
		return
	}

	// ドメイン一覧表示
	c.HTML(http.StatusOK, "domain.tmpl.html", pongo2.Context{
		"title":      config.GetEnv().AppName, // ナビゲーションメニュータイトル
		"page_title": "ドメイン一覧",
		"domainList": rows,
	})
}
