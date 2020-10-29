package user

import "C"
import (
	"errors"
	"github.com/apptut/validator"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
	"time"
	app "ucenter/src"
	"ucenter/src/models"
	"ucenter/src/utils/crypto"
)

// 主要验证密码是否正确
func InfoValid(c *gin.Context) (*models.User, error) {
	name := c.PostForm("username")
	pwd := c.PostForm("password")

	var model models.User
	err := app.Db().First(&model, "username = ? ", name).Error
	if err != nil {
		return nil, err
	}

	if !crypto.BcryptCompare(model.Password, pwd) {
		return nil, errors.New("password is error")
	}

	return &model, nil
}

// 登录完成，发送一个token，之后其他接口每次请求都需要带上这个token用于校验是否登录
// 如果token过期或者token不存在则重新提示登录
func GetToken(c *gin.Context, user *models.User) (string, error) {
	// Unix 和 UnixNano 区别是啥？
	now := strconv.FormatInt(time.Now().UnixNano(), 10)
	rand.Seed(time.Now().UnixNano())
	randStr := strconv.Itoa(rand.Intn(100000))

	uid := strconv.Itoa(int(user.ID))
	token := crypto.SHA256(uid + randStr + user.Mobile + now)

	err := app.Redis().Set(c, "login:token:"+token, uid, time.Hour*4).Err()

	if err != nil {
		return "", err
	}

	return token, nil
}

func ParamsValid(c *gin.Context) error {
	_, err := validator.New(map[string][]string{
		"username": {c.PostForm("username")},
		"password": {c.PostForm("password")},
	}, map[string]string{
		"username": "regex:^\\w{6,20}$",
		"password": "regex:^\\w{6,20}$",
	})

	if err != nil {
		return errors.New("login username or password error")
	}

	return nil
}
