package user

import (
	"github.com/gin-gonic/gin"
	"ucenter/src/service/user"
	"ucenter/src/utils"
)

func Reg(c *gin.Context) {
	// 参数验证
	err := user.Valid(c)
	if err != nil {
		utils.JsonErr(c, err)
		return
	}

	// 验证码验证
	err = user.CheckCode(c)
	if err != nil {
		utils.JsonErr(c, err)
		return
	}

	// 验证用户信息
	err = user.CheckUser(c)
	if err != nil {
		utils.JsonErr(c, err)
		return
	}

	// 保存数据
	err = user.SaveUser(c)
	if err != nil {
		utils.JsonErr(c, err)
		return
	}

	utils.JsonOK(c, "注册成功!")
}
