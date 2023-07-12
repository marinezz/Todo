<h1 align='center'>Todo</h1>

<center>
    golang简易记事本--TodoList
</center>

<p align="center">
    <a href="https://img.shields.io/badge/%E8%AF%AD%E8%A8%80-go-blue">
        <img alt="Static Badge" src="https://img.shields.io/badge/%E8%AF%AD%E8%A8%80-go-blue">
    </a>
    <a href="https://img.shields.io/badge/web%E6%A1%86%E6%9E%B6-gin-green">
        <img alt="Static Badge" src="https://img.shields.io/badge/web%E6%A1%86%E6%9E%B6-gin-green">
    </a>
    <a href="https://img.shields.io/badge/orm%E6%A1%86%E6%9E%B6-gorm-red">
        <img alt="Static Badge" src="https://img.shields.io/badge/orm%E6%A1%86%E6%9E%B6-gorm-red">
    </a>
</p>

# 🐣快速开始：

* 创建数据库：

```sql
CREATE DATABASE TodoList;
```

* 拉取项目到本地

```bash
git clone https://github.com/marinezz/Todo.git
```

* 编译并启动

```bash
go build
./Todo
```

* 访问并测试，进入浏览器输入

```bash
127.0.0.1/index
```



# 🐥项目结构：

​	项目比较简单，主要实现了常规的增、删、改、查四个操作：增加待办事项；删除代办事项；改变代办事项的状态；每次操作都会从新查询数据库最新状态

​	简单的功能还是将逻辑进行分层，其中static和templates都是前端文件，目录结构如下：

```
├─Controller // 控制层
├─dao   	 // 持久层（因为业务简单，就没有业务层）
├─models     // 模型构建
├─router	 // 路由层
├─static     // 静态文件
│  ├─css
│  ├─fonts
│  └─js
└─templates  // 模板文件 
```



# 🐤简单介绍

对于gin框架的使用，主要分为以下几步：

* 创建路由实例

```go
r := gin.Default()
```

* 加载静态文件

```go
r.Static("/static", "./static")  // 加载静态文件
r.LoadHTMLGlob("templates/*")    // 加载摹本文件
```

* 获取请求并处理（以get请求为例）

```go
r.GET("/index", controller.IndexHanlder)  // 第二个参数是处理请求的函数
```

* 启动服务（可以指定端口，默认8080）

```
r.run()
```



对于gorm的使用，简单总结为以下几步：

* 创建模型（与数据库对应的结构体，tag是对应阶段传递的json）

```go
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}
```

* 连接数据库

```go
	dsn := "root:123456@tcp(127.0.0.1:3306)/todolist?charset=utf8mb4&parseTime=True &loc=Local"
	DB, _ = gorm.Open("mysql", dsn)
```

* 绑定模型，如果数据库不存在会自动创建

```go
dao.DB.AutoMigrate(&models.Todo{})
```

* 关闭数据库

```go
DB.close()
```



# 🐔参考手册

gin：[Gin Web Framework]([Gin Web Framework (gin-gonic.com)](https://gin-gonic.com/zh-cn/))

gorm：[GORM]([GORM 指南 | GORM - The fantastic ORM library for Golang, aims to be developer friendly.](https://gorm.io/zh_CN/docs/index.html))
