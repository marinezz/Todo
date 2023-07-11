package models

import "Todo/dao"

// Todo 创建模型 (属性分别为id,名称，状态)
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

/*
	todo这个层所有的增删改查操作都在这里
*/

// CreateATodo 创建todo
func CreateATodo(todo *Todo) (err error) {
	if err = dao.DB.Create(&todo).Error; err != nil {
		return err
	}
	return
}

// GetAllTodo 查询所有数据
func GetAllTodo() (todos []*Todo, err error) {
	if err := dao.DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

// GetATodo 查询一个数据
func GetATodo(id string) (t1 *Todo, err error) {
	// 返回值定义的变量不会设置默认值，会出现野指针情况，需要先初始化
	t1 = new(Todo)
	if err := dao.DB.Where("id=?", id).First(&t1).Error; err != nil {
		return nil, err
	}
	return t1, nil
}

// SaveTodo 保存数据
func SaveTodo(todo *Todo) (err error) {
	if err = dao.DB.Save(todo).Error; err != nil {
		return err
	}
	return
}

// Delete 删除数据
func Delete(id string) (err error) {
	err = dao.DB.Delete(Todo{}, id).Error
	return
}
