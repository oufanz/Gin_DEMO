package auth

import (
	"Gin_DEMO/config"
	"Gin_DEMO/utils/jwtutil"
	"Gin_DEMO/utils/logs"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(r *gin.Context) {
	//1.获取前端用户名的密码
	UserInfo := UserInfo{}
	if err := r.ShouldBindJSON(&UserInfo); err != nil {
		r.JSON(200, gin.H{
			"message": err.Error(),
			"status":  401,
		})
		return
	}
	logs.Debug(map[string]interface{}{
		"用户名": UserInfo.Username,
		"密码":  UserInfo.Password,
	}, "开始验证登录信息")
	if UserInfo.Username == config.Username && UserInfo.Password == config.Password {
		//认证成功
		//生成JWT的token
		ss, err := jwtutil.GenToken(UserInfo.Username)
		if err != nil {
			logs.Error(map[string]interface{}{
				"用户名": UserInfo.Username,
			}, "登录密码正确生成token失败")
			r.JSON(200, gin.H{
				"message": "生成token失败",
				"status":  401,
			})
			return
		} else {
			logs.Info(map[string]interface{}{"用户名": UserInfo.Username}, "登录成功")
			data := make(map[string]interface{})
			data["token"] = ss
			r.JSON(200, gin.H{
				"status":  200,
				"message": "登录成功",
				"data":    data,
			})
		}
	} else {
		logs.Error(map[string]interface{}{
			"正确用户名为": config.Username,
			"正确密码为":  config.Password,
		}, "登录失败")
		r.JSON(200, gin.H{
			"message": "用户名或者密码错误",
			"status":  200,
		})
	}

}
func Logout(r *gin.Context) {
	//退出
	//实现退出逻辑
	r.JSON(200, gin.H{
		"message": "退出成功",
		"status":  200,
	})
	logs.Debug(nil, "退出成功")
}
