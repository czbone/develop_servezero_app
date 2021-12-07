package filters

import (
	"web/config"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RegisterSession() gin.HandlerFunc {
	store, _ := sessions.NewRedisStore(
		10,
		"tcp",
		config.GetEnv().RedisIp+":"+config.GetEnv().RedisPort,
		config.GetEnv().RedisPassword,
		[]byte(config.GetEnv().SessionSecret))
	return sessions.Sessions(config.GetEnv().SessionKey, store)
}
