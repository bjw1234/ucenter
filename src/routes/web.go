// 请求的url配置文件
package routes

import (
	"github.com/gin-gonic/gin"
	"ucenter/src/controllers/sms"
	"ucenter/src/controllers/user"
	"ucenter/src/middleware"
)

func New() *gin.Engine {
	engine := gin.New()

	engine.GET("/v1/sms/send", sms.Send)
	engine.POST("/v1/user/reg", user.Reg)
	engine.POST("/v1/user/login", user.Login)
	engine.GET("/v1/user/profile",middleware.Auth(), user.Profile)

	return engine
}
