package routes

import (
	api "gin_mall/api/v1" // 自定义命名导入
	"gin_mall/middleware"

	"github.com/gin-gonic/gin"
)

// 项目逻辑:路由调用api(控制器函数)，api函数中再调用对应的service服务

//路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors()) // 开启跨域

	v1 := r.Group("api/v1")
	{
		//用户登录和注册
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
	}

	return r
}
