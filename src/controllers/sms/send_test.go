package sms

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-redis/redis/v8"
	"path"
	"runtime"
	"testing"
	"time"
	app "ucenter/src"
	"ucenter/src/service/sms"
)

func TestSend(t *testing.T) {
	code := sms.PostCode()
	t.Log("随机验证码为：", code)
}

func TestRedis(t *testing.T) {
	var c = context.Background()
	app.Redis().Set(c, "name", "bjw", time.Minute)
	strCmd := app.Redis().Get(c, "name")
	flag := strCmd.Err() == redis.Nil
	t.Log(flag)
	t.Log(strCmd.Err().Error())
}

func TestFlag(t *testing.T) {
	mode := flag.Bool("d", false, "debug mode")
	flag.Parse()

	_, filePath, _, _ := runtime.Caller(0)
	str := path.Dir(filePath)
	dir := path.Base(filePath)
	fmt.Println("file:", str)
	fmt.Println("dir: ", dir)
	fmt.Println("mode: ", *mode)
}
