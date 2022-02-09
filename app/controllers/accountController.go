package controllers

import (
	"net/http"
	"strings"
	"web/config"
	"web/db"
	"web/modules/filters/auth"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ######################################################################
// ページコントローラ
// ・GETまたはPOSTでHTTPリクエストを受信し、WEB画面を作成して、HTMLデータを返す
// ######################################################################
type AccountController struct{}

type ValidateUser struct {
	Account         string `validate:"required,email"`
	Password        string `validate:"required"`
	PasswordNew     string `validate:"required,max=10,min=6"`
	PasswordConfirm string `validate:"required"`
}

func (pc *AccountController) Index(c *gin.Context) {
	// パラメータ初期化
	var error, success string // メッセージパラメータ
	userDb := &db.UserDb{}

	// 入力値取得
	var account, password, passwordNew, passwordConfirm string
	act := strings.TrimSpace(c.PostForm("act")) // 実行操作

	if act == "update" {
		// 入力値取得
		account = strings.TrimSpace(c.PostForm("account"))
		password = c.PostForm("password")                // 現在のパスワード
		passwordNew = c.PostForm("password_new")         // 新規パスワード
		passwordConfirm = c.PostForm("password_confirm") // 新規パスワード(確認)

		// 入力値チェック
		newUser := &ValidateUser{
			Account:         account,
			Password:        password,
			PasswordNew:     passwordNew,
			PasswordConfirm: passwordConfirm,
		}
		validate := validator.New()
		err := validate.Struct(newUser)
		if err != nil {
			// エラーメッセージ設定
			for _, err := range err.(validator.ValidationErrors) {
				fieldName := err.Field() // 入力エラーの構造体変数名を取得
				switch fieldName {
				case "Account":
					switch err.Tag() {
					case "required":
						error = "Eメールを入力してください"
					case "email":
						error = "Eメールのフォーマットが不正です"
					}
				case "Password":
					switch err.Tag() {
					default:
						error = "現在のパスワードを入力してください"
					}
				case "PasswordNew":
					switch err.Tag() {
					case "required":
						error = "新しいパスワードを入力してください"
					default:
						error = "新しいパスワードは6文字以上10文字以下で入力してください"
					}
				case "PasswordConfirm":
					switch err.Tag() {
					default:
						error = "新しいパスワード(再入力)を入力してください"
					}
				}
				if error != "" {
					break
				}
			}
		}
	}
	// ユーザ情報取得
	//var account string
	userId := int(auth.GetDefaultUser(c)["id"].(int64))
	//userId := auth.GetDefaultUser(c)["id"].(int)
	row := userDb.GetUser(userId)
	if row == nil { // ユーザが存在しないとき
		error = "ユーザが見つかりません。"
	} else {
		account = row["account"].(string)
	}

	if error == "" { // エラーなしの場合
		// ドメイン一覧表示
		c.HTML(http.StatusOK, "account.tmpl.html", pongo2.Context{
			"app_name":         config.GetEnv().AppName, // ナビゲーションメニュータイトル
			"page_title":       "アカウント",
			"account":          account,
			"password":         password,
			"password_new":     passwordNew,
			"password_confirm": passwordConfirm,
			"success":          success,
		})
	} else {
		// ドメイン一覧表示
		c.HTML(http.StatusOK, "account.tmpl.html", pongo2.Context{
			"app_name":         config.GetEnv().AppName, // ナビゲーションメニュータイトル
			"page_title":       "アカウント",
			"account":          account,
			"password":         password,
			"password_new":     passwordNew,
			"password_confirm": passwordConfirm,
			"error":            error,
		})
	}
}
