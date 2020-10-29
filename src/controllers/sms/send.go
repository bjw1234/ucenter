package sms

import (
	"errors"
	"fmt"
	"github.com/apptut/validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"ucenter/src/service/sms"
	"ucenter/src/utils"
)

func Send(c *gin.Context) {
	ip := c.ClientIP()

	mobile, err := validateSendParams(c)

	fmt.Println("ip地址为：", ip)
	fmt.Println("手机号为：", mobile)
	if err != nil {
		utils.JsonErr(c, err)
		return
	}

	// 获取频率验证
	if !sms.SendEnable(c, ip, mobile) {
		utils.JsonErr(c, errors.New("发送频率过快"))
		return
	}

	// 发送验证码
	code := sms.PostCode(c, mobile)

	// 更新
	sms.UpdateSmsCache(c, ip, mobile)

	c.JSON(http.StatusOK, gin.H{
		"code":      0,
		"message":   "短信验证码已发送",
		"post_code": code,
	})
}

// 校验手机号是否合法
func validateSendParams(c *gin.Context) (string, error) {
	mobile := c.Query("mobile")

	_, err := validator.New(map[string][]string{
		"mobile": {mobile},
	}, map[string]string{
		"mobile": "mobile",
	}, map[string]string{
		"mobile": "手机号码格式不正确！",
	})

	if err != nil {
		return mobile, err
	}

	return mobile, nil
}
