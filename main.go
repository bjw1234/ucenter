/**
应用启动流程
main.go -> routes -> middles  -> controllerA、controllerB
*/

package main

import (
	"flag"
	"fmt"
	"path"
	"runtime"
	app "ucenter/src"
	"ucenter/src/routes"
)

func parseArgs() string {
	prjHome := flag.String("prjHome", "", "project home dir")
	flag.Parse()

	if *prjHome == "" { // 默认为空
		_, filePath, _, ok := runtime.Caller(0)
		if ok {
			return path.Dir(filePath)
		}
		fmt.Println("params error!")
	}

	return *prjHome
}

func main() {
	// 解析参数
	prjHome := parseArgs()
	fmt.Println("prjHome: ", prjHome)
	// 初始化配置
	err := app.Init(prjHome)
	if err != nil {
		app.Destruct()
	}

	// 启动
	engine := routes.New()
	err = engine.Run(":9527")
	if err != nil {
		fmt.Println(err)
	}
}
