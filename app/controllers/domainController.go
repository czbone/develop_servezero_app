package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"web/config"
	"web/db"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"

	"github.com/go-playground/validator/v10"
	"github.com/lithammer/shortuuid"
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
			// エラーメッセージ設定
			for _, err := range err.(validator.ValidationErrors) {
				fieldName := err.Field() // 入力エラーの構造体変数名を取得

				switch fieldName {
				case "Name":
					switch err.Tag() {
					case "fqdn":
						error = "ドメイン名のフォーマットが不正です"
					}
				}
			}
		}

		if error == "" {
			// ドメイン存在確認
			row := domainDb.GetDomainByName(name)
			if row == nil {
				// ドメインID作成

				// ドメイン追加
				domainId := generateDomainId(name)
				result := domainDb.AddDomain(name, "abc", domainId)

				if result {
					success = "ドメインを登録しました"
				} else {
					error = "ドメイン登録に失敗しました"
				}
			} else {
				error = "登録済みのドメインです"
			}
		}
	} else if act == "del" { // ドメイン削除の場合
		// 入力値取得
		domainId, _ := strconv.Atoi(strings.TrimSpace(c.PostForm("id"))) // ドメインID

		// ドメイン情報取得
		row := domainDb.GetDomain(domainId)
		if row == nil {
			error = "ドメイン削除に失敗しました"
		} else {
			domainName := row["name"]

			// ドメイン削除
			result := domainDb.DelDomain(domainId)
			if result {
				success = fmt.Sprintf("ドメイン「%s」を削除しました", domainName)
			} else {
				error = "ドメイン削除に失敗しました"
			}
		}
	}

	// ドメイン一覧取得
	rows := domainDb.GetDomainList()

	if error == "" { // エラーなしの場合
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

func generateDomainId(domain string) string {
	domainId := shortuuid.NewWithNamespace(domain + time.Now().Format("2006-01-02 15:04:05"))
	return domainId
}
