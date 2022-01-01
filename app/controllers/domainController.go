package controllers

import (
	"net/http"
	"strconv"
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
	// パラメータ初期化
	var error, success string // メッセージパラメータ
	domainDb := &db.DomainDb{}

	// 入力値取得
	act := strings.TrimSpace(c.PostForm("act")) // 実行操作

	if act == "add" { // ドメイン追加の場合
		// 入力値取得
		name := strings.TrimSpace(c.PostForm("name")) // ドメイン名

		// 入力値チェック
		newDomain := &ValidateNewDomain{
			Name: name,
		}
		validate := validator.New()
		err := validate.Struct(newDomain)
		if err != nil {
			error = "不正なドメイン名です"
		}

		// ドメイン存在確認
		row := domainDb.GetDomainByName(name)
		if row == nil {
			// ドメイン追加
			result := domainDb.AddDomain(name, "abc")

			if result {
				success = "ドメイン登録しました"
				/*c.HTML(http.StatusOK, "domain.tmpl.html", pongo2.Context{
					"title":      config.GetEnv().AppName, // ナビゲーションメニュータイトル
					"page_title": "ドメイン一覧",
					"domainList": rows,
					"success":    "ドメイン登録しました",
				})*/
			} else {
				error = "ドメイン登録に失敗しました"
				/*c.HTML(http.StatusOK, "domain.tmpl.html", pongo2.Context{
					"title":      config.GetEnv().AppName, // ナビゲーションメニュータイトル
					"page_title": "ドメイン一覧",
					"domainList": rows,
					"error":      "ドメイン登録に失敗しました",
				})*/
			}
		} else {
			error = "登録済みのドメインです"
			/*c.HTML(http.StatusOK, "domain.tmpl.html", pongo2.Context{
				"title":      config.GetEnv().AppName, // ナビゲーションメニュータイトル
				"page_title": "ドメイン一覧",
				"domainList": rows,
				"error":      "登録済みのドメインです",
			})*/
		}
	} else if act == "del" { // ドメイン削除の場合
		// 入力値取得
		domainId, _ := strconv.Atoi(strings.TrimSpace(c.PostForm("id"))) // ドメインID

		// ドメイン情報取得
		row := domainDb.GetDomain(domainId)
		if row != nil {
			error = "ドメイン登録に失敗しました"
		}

	}

	// ドメイン一覧取得
	rows := domainDb.GetDomainList()

	if error == "" {
		// ドメイン一覧表示
		c.HTML(http.StatusOK, "domain.tmpl.html", pongo2.Context{
			"title":      config.GetEnv().AppName, // ナビゲーションメニュータイトル
			"page_title": "ドメイン一覧",
			"domainList": rows,
			"success":    success,
		})
	} else {
		c.HTML(http.StatusOK, "domain.tmpl.html", pongo2.Context{
			"title":      config.GetEnv().AppName, // ナビゲーションメニュータイトル
			"page_title": "ドメイン一覧",
			"domainList": rows,
			"error":      error,
		})
	}
}
