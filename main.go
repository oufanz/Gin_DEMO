package main

import (
	"Gin_DEMO/config"
	_ "Gin_DEMO/config" //导入包但不使用里面的方法
	"Gin_DEMO/utils/jwtutil"
	"Gin_DEMO/utils/logs"
	"fmt"

	"github.com/gin-gonic/gin"
)

func test_jwt_token() {
	ss, _ := jwtutil.GenToken("ddd")
	fmt.Println(ss)
}

func main() {
	//1.加载程序配置:有方法叫做initial，自动调用
	logs.Info(nil, "程序启动！")
	r := gin.Default() //1.初始化，创建一个路由引擎，可以
	r.Run(config.Port) //阻塞
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
