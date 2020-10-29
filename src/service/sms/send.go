package sms

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"strconv"
	"time"
	app "ucenter/src"
	"ucenter/src/utils/crypto"
)

func UpdateSmsCache(c *gin.Context, ip string, phone string) {
	redisKey := "sms:" + crypto.MD5(ip+phone)
	app.Redis().Set(c, redisKey, true, time.Minute)
}

func SendEnable(c *gin.Context, ip string, phone string) bool {
	redisKey := "sms:" + crypto.MD5(ip+phone)
	fmt.Println("redisKey", redisKey)
	redisErr := app.Redis().Get(c, redisKey).Err()
	if redisErr != nil {
		return redisErr == redis.Nil
	}

	return false
}

func PostCode(c *gin.Context, mobile string) string {
	var code string

	rand.Seed(time.Now().Unix())

	for i := 1; i < 5; i++ {
		code += strconv.Itoa(rand.Intn(10))
	}

	// 为后续注册做准备
	app.Redis().Set(c, "reg:"+mobile, code, time.Minute*30)

	return code
}
