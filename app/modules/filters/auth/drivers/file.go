package drivers

import (
	"net/http"
	"os"
	"web/config"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

var store *sessions.FilesystemStore

type fileAuthManager struct {
	name string
}

func init() {
	store = sessions.NewFilesystemStore(os.TempDir(), securecookie.GenerateRandomKey(32), securecookie.GenerateRandomKey(32))

	store.Options = &sessions.Options{
		// Domain:   "localhost", // Set this when we have a domain name.
		Path:     "/",
		MaxAge:   3600 * 2, // 有効期間2時間
		HttpOnly: true,
		// Secure:   true, // Set this when TLS is set up.
	}
}

func NewFileAuthDriver() *fileAuthManager {
	return &fileAuthManager{
		name: config.GetCookieConfig().NAME,
	}
}

func (fileAuth *fileAuthManager) Check(c *gin.Context) bool {
	session, err := store.Get(c.Request, fileAuth.name)
	if err != nil {
		return false
	}
	if session == nil {
		return false
	}
	if session.Values == nil {
		return false
	}
	if session.Values["id"] == nil {
		return false
	}
	return true
}

func (fileAuth *fileAuthManager) User(c *gin.Context) interface{} {
	session, err := store.Get(c.Request, fileAuth.name)
	if session == nil {
		panic("wrong session")
	}
	if err != nil {
		return session.Values
	}
	return session.Values
}

func (fileAuth *fileAuthManager) Login(http *http.Request, w http.ResponseWriter, user map[string]interface{}) interface{} {
	session, err := store.Get(http, fileAuth.name)
	if err != nil {
		return false
	}
	session.Values["id"] = user["id"]
	_ = session.Save(http, w)
	return true
}

func (fileAuth *fileAuthManager) Logout(http *http.Request, w http.ResponseWriter) bool {
	session, err := store.Get(http, fileAuth.name)
	if err != nil {
		return false
	}
	session.Values["id"] = nil
	_ = session.Save(http, w)
	return true
}
