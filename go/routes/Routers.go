package routes

import (
	"github.com/gin-gonic/gin"
	"task0/go/common/util"
	"task0/go/controller"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	// 用户User路由组，类似于Java的RequestMapping
	userGroup := r.Group("/users")
	{
		// 增加用户User
		userGroup.POST("", controller.CreateUser)
		// 查看所有的User
		userGroup.GET("", util.JWTAuthMiddleware(), controller.GetUserList)
		// 修改某个User
		userGroup.PUT("", util.JWTAuthMiddleware(), controller.UpdateUser)
		// 删除某个User
		userGroup.DELETE("", util.JWTAuthMiddleware(), controller.DeleteUser)
	}

	authGroup := r.Group("")
	{
		// 鉴权
		authGroup.GET("/login", controller.AuthHandler)
	}

	return r
}
