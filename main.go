package main

import "github.com/gin-gonic/gin"

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{ // 返回一个Json的数据，状态码为200
		"message": "hello golang",
	})
}

func main() {
	r := gin.Default() // 返回默认的路由引擎

	// 指定用户使用get请求访问hello，执行sayHello函数
	r.GET("hello", sayHello)

	// 启动服务,括号中可以指定端口
	r.Run()
}
