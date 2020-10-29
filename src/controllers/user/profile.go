package user

import (
	"github.com/gin-gonic/gin"
	app "ucenter/src"
	"ucenter/src/models"
	"ucenter/src/utils"
)

func Profile(c *gin.Context) {
	uid := c.GetString("uid")
	var user models.User

	err := app.Db().First(&user, uid).Error
	if err != nil {
		utils.JsonErr(c, err)
		return
	}

	c.JSON(0, gin.H{
		"username":   user.Username,
		"mobile":     user.Mobile,
		"created_at": user.CreatedAt.Unix(),
	})
}
