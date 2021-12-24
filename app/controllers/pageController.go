package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"web/db"
	"web/modules/filters/auth"
	"web/modules/log"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// ######################################################################
// ページコントローラ
// ・GETまたはPOSTでHTTPリクエストを受信し、WEB画面を作成して、HTMLデータを返す
// ######################################################################
type PageController struct{}

func (pc *PageController) Index(c *gin.Context) {

	log.Println("#index controller")

	GetAllData(c)
}

func GetAllData(c *gin.Context) {
	posts := []string{
		"Larry Ellison",
		"Carlos Slim Helu",
		"Mark Zuckerberg",
		"Amancio Ortega ",
		"Jeff Bezos",
		" Warren Buffet ",
		"Bill Gates",
		"selman tunc",
	}
	// Call the HTML method of the Context to render a template
	c.HTML(http.StatusOK, "index.tmpl.html", pongo2.Context{
		"title": "hello pongo",
		"posts": posts,
	})
}

func (pc *PageController) Login(c *gin.Context) {
	account := strings.TrimSpace(c.PostForm("account"))
	password := c.PostForm("password")

	// ユーザ情報を取得
	userDb := &db.UserDb{}
	row := userDb.GetUser(account)
	if row != nil {
		authChecked := false

		// パスワードチェック
		passByte := []byte(fmt.Sprintf("%v", row["password"]))
		err := bcrypt.CompareHashAndPassword(passByte, []byte(password)) // cost=10
		if err == nil {                                                  // パスワードチェックOKの場合
			authChecked = true
		}

		if authChecked {
			// ### ユーザ認証に成功 ###
			// セッションにログイン情報を登録
			authDr, _ := c.MustGet("web_auth").(auth.Auth)
			authDr.Login(c.Request, c.Writer, map[string]interface{}{"id": row["id"]})

			c.Redirect(http.StatusFound, "/")
			return
		}
	}

	c.HTML(http.StatusOK, "login.tmpl.html", pongo2.Context{
		"error": "ログインに失敗しました",
	})
}

func (pc *PageController) Logout(c *gin.Context) {
	// セッションからログイン情報を削除
	authDr, _ := c.MustGet("web_auth").(auth.Auth)
	authDr.Logout(c.Request, c.Writer)

	c.Redirect(http.StatusFound, "/")
}
