package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

// InitMysql 连接数据库
func InitMysql() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local"
	DB, _ = gorm.Open("mysql", dsn)
	return
}

// CloseMysql 关闭数据库
func CloseMysql() (err error) {
	DB.Close()
	return
}
