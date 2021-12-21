package controllers

import (
	"net/http"
	"web/db"
	"web/modules/filters/auth"
	"web/modules/log"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

// ######################################################################
// ページコントローラ
// ・GETまたはPOSTでHTTPリクエストを受信し、WEB画面を作成して、HTMLデータを返す
// ######################################################################
type PageController struct{}

func (pc *PageController) Index(c *gin.Context) {

	log.Println("#index controller")
	/*
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "GO GO GO!",
		})*/
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
	log.Println("Login controller....")

	account := c.PostForm("account")
	pass := c.PostForm("password")

	log.Println("login:" + account + "- " + pass)

	authDr, _ := c.MustGet("web_auth").(auth.Auth)

	userDb := &db.UserDb{}
	row := userDb.GetUser(account)
	if row == nil {
		c.HTML(http.StatusOK, "login.tmpl.html", pongo2.Context{
			"error": "ログインに失敗しました",
		})
	} else {
		// セッションにログイン情報を登録
		authDr.Login(c.Request, c.Writer, map[string]interface{}{"id": row["id"]})

		c.Redirect(http.StatusFound, "/")
	}
}
