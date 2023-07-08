package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{ // 返回一个Json的数据，状态码为200
		"message": "hello golang",
	})
}

func main() {
	r := gin.Default() // 返回默认的路由引擎

	// 指定用户使用get请求访问hello，执行sayHello函数
	r.GET("hello", sayHello)

	// rest风格请求
	r.GET("book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "Get",
		})
	})

	r.POST("book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "Post",
		})
	})

	r.DELETE("book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "Delete",
		})
	})

	r.PUT("book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "Put",
		})
	})

	// 启动服务,括号中可以指定端口
	r.Run()
}
