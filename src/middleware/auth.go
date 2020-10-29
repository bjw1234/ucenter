package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
	app "ucenter/src"
	"ucenter/src/utils"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authVal := c.GetHeader("Authorization")
		uid, err := checkAuth(c, authVal)
		if err != nil {
			utils.JsonErr(c, err)
			c.Abort() // ?
			return
		}

		// 为什么要set？ 设置完成后，就可以在后续拿到当前用户信息了
		c.Set("uid", uid)

		c.Next()
	}
}

func checkAuth(c *gin.Context, authVal string) (string, error) {
	if len(authVal) <= 0 {
		return "", errors.New("auth is not exist")
	}

	arr := strings.Split(authVal, " ")
	if len(arr) != 2 {
		return "", errors.New("auth params error")
	}

	token := strings.TrimSpace(arr[1])
	uid, err := app.Redis().Get(c, "login:token:"+token).Result()
	if err != nil {
		return "", errors.New("current user auth failed")
	}

	return uid, nil
}
