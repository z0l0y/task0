package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"task0/go/dao"
	"task0/go/routes"
)

func main() {
	// 连接postgres数据库
	err := dao.InitPostgreSQL()
	if err != nil {
		panic(err)
	}
	// 程序退出关闭数据库连接
	defer dao.Close()
	// 注册路由
	r := routes.SetRouter()
	// 启动端口为8081的项目
	r.Run(":8081")
}
