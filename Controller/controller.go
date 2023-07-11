package controller

import (
	"Todo/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHanlder(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodo(c *gin.Context) {
	// 取出请求中的数据
	var t1 models.Todo
	_ = c.ShouldBind(&t1)

	// 存入数据库，返回响应
	err := models.CreateATodo(&t1)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, t1)
	}
}

func GetTodo(c *gin.Context) {
	// 查询所有的数据
	todos, err := models.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	} else {
		// 取到正确的结果
		c.JSON(http.StatusOK, todos)
	}
}

func UpdataTodo(c *gin.Context) {
	// 获取要更新的id
	id := c.Param("id")

	// 根据id更新状态,先查询，再修改
	todo, err := models.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	}
	// 和Json数据绑定，在这里更新了状态
	c.BindJSON(&todo)
	// 保存进入数据库
	if err = models.SaveTodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteTodo(c *gin.Context) {
	// 获取要删除的id
	id := c.Param("id")

	// 数据库中删除
	models.Delete(id)
}
