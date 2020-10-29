package user

import (
	"github.com/gin-gonic/gin"
	"ucenter/src/service/user"
	"ucenter/src/utils"
)

func Login(c *gin.Context) {
	err := user.ParamsValid(c)
	if err != nil {
		utils.JsonErr(c, err)
		return
	}

	currentUser, err := user.InfoValid(c)
	if err != nil {
		utils.JsonErr(c, err)
		return
	}

	token, err := user.GetToken(c, currentUser)
	if err != nil {
		utils.JsonErr(c, err)
		return
	}

	utils.JsonOK(c, token)
}
