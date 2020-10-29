package user

import (
	"errors"
	"github.com/apptut/validator"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	app "ucenter/src"
	"ucenter/src/models"
	"ucenter/src/utils/crypto"
)

func SaveUser(c *gin.Context) error {
	pwd := c.PostForm("password")
	bcryptPwd, err := crypto.BcryptPwd(pwd)

	user := models.User{
		Username: c.PostForm("username"),
		Mobile:   c.PostForm("mobile"),
		Password: bcryptPwd,
	}

	err = app.Db().Create(&user).Error

	if err != nil {
		return errors.New("bcrypt pwd or create user error")
	}

	return nil
}

func CheckUser(c *gin.Context) error {
	username := c.PostForm("username")
	mobile := c.PostForm("mobile")

	var model models.User
	err := app.Db().Find(&model, "username = ? or mobile = ?", username, mobile).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
	}

	return errors.New("user or mobile already exist")
}

func CheckCode(c *gin.Context) error {
	mobile := c.PostForm("mobile")
	code := c.PostForm("code")

	redisKey := "reg:" + mobile
	savedCode, err := app.Redis().Get(c, redisKey).Result()

	if err != nil {
		return errors.New("please get code first")
	}

	if code != savedCode {
		return errors.New("code is error")
	}

	return nil
}

func Valid(c *gin.Context) error {
	_, err := validator.New(map[string][]string{
		"username": {c.PostForm("username")},
		"password": {c.PostForm("password")},
		"mobile":   {c.PostForm("mobile")},
		"code":     {c.PostForm("code")},
	}, map[string]string{
		"username": "regex:^[\\w_]{6,20}$",
		"mobile":   "mobile",
		"password": "regex:^[\\w]{6,20}",
		"code":     "regex:^[\\d]{4}$",
	})

	if err != nil {
		return err
	}

	return nil
}
