package router

import (
	controller "Todo/Controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// 创建路由
	// 创建一个路由的实例（返回默认路由引擎）
	r := gin.Default()
	// 加载静态文件(两个参数：相对路径，实际的路径)
	r.Static("/static", "./static")
	// 告诉Gin在哪里去找页面文件
	r.LoadHTMLGlob("templates/*")

	r.GET("/index", controller.IndexHanlder)

	// 创建路由组
	v1Group := r.Group("v1")
	{
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)

		// 查看所有代办事项,每次操作都会发送该请求
		v1Group.GET("/todo", controller.GetTodo)

		// 修改
		v1Group.PUT("/todo/:id", controller.UpdataTodo)

		// 删除
		v1Group.DELETE("todo/:id", controller.DeleteTodo)
	}
	return r
}
