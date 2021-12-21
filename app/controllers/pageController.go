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
	//id := c.Param("userid")
	id := 123

	//rs := db.Query("select name,avatar,id from users where id = ?", id)

	userDb := &db.UserDb{}
	//userDb.GetUser(account)
	userDb.Test(account)

	// セッションにログイン情報を登録
	authDr.Login(c.Request, c.Writer, map[string]interface{}{"id": id})

	c.HTML(http.StatusOK, "login.tmpl.html", pongo2.Context{
		"title": "GO Login!!!",
		"error": "GO Login!!!",
	})
}

/*
func DBExample(c *gin.Context) {

	// 数据库插入
	insertRs, _ := db.Exec("insert into users (name, avatar, sex) values (?, ?, ?)", "人才", "unknown", 1)
	insertId, _ := insertRs.LastInsertId()
	log.Printf("insert id: %d\n", insertId)

	// 数据库更新
	db.Exec("update users set name = ? where id = ?", "饭桶", insertId)

	// 数据库中间件
	_, _ = database.Table("users").Where("id", "=", insertId).Update(database.H{
		"name": "你好",
	})

	// 数据库查询
	rs := db.Query("select name,avatar,id from users where id < ?", 100)
	log.Println(rs[0]["name"])

	rs1, _ := database.Table("users").
		Select("name", "avatar", "id").
		Where("id", "<", 100).
		All()
	log.Println(rs1[0])

	// 数据库事务
	_, _ = db.WithTransaction(func(tx *db.SqlTxStruct) (error, map[string]interface{}) {
		_, err := tx.Query("select name,avatar,id from users where id < ?", 100)
		if err != nil {
			return err, map[string]interface{}{}
		}
		return nil, map[string]interface{}{}
	})

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{
			"query_result": rs,
		},
	})
}
*/
