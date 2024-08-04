package main

import (
	_ "Gin_DEMO/config" //导入包但不使用里面的方法
	"Gin_DEMO/utils/logs"

	"github.com/gin-gonic/gin"
)

func main() {
	//1.加载程序配置:有方法叫做initial，自动调用
	logs.Info(nil, "程序启动！")
	r := gin.Default() //1.初始化，创建一个路由引擎，可以
	r.Run()            //阻塞
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
