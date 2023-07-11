package main

import (
	"Todo/dao"
	"Todo/models"
	"Todo/router"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// 创建数据库 sql : CREATE DATABASE TodoList;
	// 连接数据库
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}
	defer dao.CloseMysql()

	// 绑定模型
	dao.DB.AutoMigrate(&models.Todo{})

	// 路由
	r := router.SetupRouter()

	// 启动服务(括号中可以指定端口，默认8080)
	r.Run()
}
