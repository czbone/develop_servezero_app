package auth

import (
	"net/http"
	"web/modules/filters/auth/drivers"
	"web/modules/log"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

const (
	FileAuthDriverKey = "file"
)

var driverList = map[string]Auth{
	FileAuthDriverKey: drivers.NewFileAuthDriver(),
}

type Auth interface {
	Check(c *gin.Context) bool
	User(c *gin.Context) interface{}
	Login(http *http.Request, w http.ResponseWriter, user map[string]interface{}) interface{}
	Logout(http *http.Request, w http.ResponseWriter) bool
}

func RegisterGlobalAuthDriver(authKey string, key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(key, GenerateAuthDriver(authKey))
		c.Next()
	}
}

func Middleware(authKey string) gin.HandlerFunc {
	return func(c *gin.Context) {

		name := c.PostForm("account")
		pass := c.PostForm("password")

		log.Println("login:" + name + "- " + pass)
		if !GenerateAuthDriver(authKey).Check(c) {
			c.HTML(http.StatusOK, "login.tmpl.html", pongo2.Context{
				"title": "login first",
			})
			c.Abort()
		}
		c.Next()
	}
}

func GenerateAuthDriver(string string) Auth {
	auth := driverList[string]
	if auth == nil {
		log.Errorf("Auth driver not found: %v", string)
	}
	return auth
}

func GetCurUser(c *gin.Context, key string) map[string]interface{} {
	authDriver, _ := c.MustGet(key).(Auth)
	return authDriver.User(c).(map[string]interface{})
}
