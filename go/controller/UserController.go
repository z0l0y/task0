package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task0/go/common/rsp"
	"task0/go/common/util"
	"task0/go/entity"
	"task0/go/service"
)

// *gin.Context，代表了当前 HTTP 请求的上下文信息，包括请求参数、请求头、响应

func CreateUser(c *gin.Context) {
	// 定义一个User变量
	var user entity.User
	// 将request请求中的body数据根据json格式解析到User结构变量中
	// 这里类似于我们Java里面的@RequestBody
	errBind := c.BindJSON(&user)
	if errBind != nil {
		return
	}
	// 将user传给service层，这里和Java不同，没有ServiceImpl
	err := service.CreateUser(&user)
	// 判断是否异常
	if err != nil {
		rsp.Error(c, err.Error())
	} else {
		rsp.Success(c, "新增成功", user)
	}
}

func GetUserList(c *gin.Context) {
	todoList, err := service.GetAllUser()
	if err != nil {
		rsp.Error(c, err.Error())
	} else {
		rsp.Success(c, "请求成功", todoList)
	}
}

func UpdateUser(c *gin.Context) {
	username, ok := c.Get("username")
	if !ok {
		rsp.Error(c, "无效的username")
	}
	user, err := service.GetUserByName(username.(string))
	if err != nil {
		rsp.Error(c, err.Error())
		return
	}
	if user == nil {
		rsp.Error(c, "找不到用户")
		return
	}
	errBind := c.BindJSON(&user)
	if errBind != nil {
		return
	}
	if err = service.UpdateUser(user); err != nil {
		rsp.Error(c, err.Error())
	} else {
		rsp.Success(c, "更新成功", user)
	}
}

func DeleteUser(c *gin.Context) {
	username, ok := c.Get("username")
	if !ok {
		rsp.Error(c, "无效的username")
	}
	if err := service.DeleteUserByUsername(username.(string)); err != nil {
		rsp.Error(c, err.Error())
	} else {
		rsp.Success(c, "删除成功", username)
	}
}

func AuthHandler(c *gin.Context) {
	// 用户发送用户名和密码过来
	var user *entity.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "无效的参数",
		})
		return
	}
	errQuery := service.GetUser(user)
	if errQuery == nil {
		// 生成Token
		tokenString, _ := util.GenToken(user.Name)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{"token": tokenString},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 400,
		"msg":  "鉴权失败,未找到匹配的用户",
	})
	return
}
