package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ######################################################################
// APIコントローラ
// ・JSONデータを受信し、JSONデータを返す
// ######################################################################
type ApiController struct{}

func (ac *ApiController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 404,
		"msg":  "ページが見つかりません",
	})
}
