package auth

import (
	"Gin_DEMO/controllers/auth"

	"github.com/gin-gonic/gin"
)

func login(g *gin.RouterGroup) {
	g.POST("/login", auth.Login)
}
func logout(g *gin.RouterGroup) {
	g.POST("/logout", auth.Logout)

}
func RegisterSubRouters(g *gin.RouterGroup) {
	//配置登录功能路由策略
	authGroup := g.Group("/auth")
	login(authGroup)
	logout(authGroup)
	//登录:
}
