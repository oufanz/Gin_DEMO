// 项目路由信息
package rooters

import (
	"Gin_DEMO/rooters/auth"

	"github.com/gin-gonic/gin"
)

func RegisterRouters(r *gin.Engine) {
	//登录的路由配置
	//1.登录：login
	//2.退出
	apiGroup := r.Group("/api")
	auth.RegisterSubRouters(apiGroup)
}
