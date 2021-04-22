package model

import (
	"github.com/tzr2020/gin_demo/dao"
)

// Todo Model
type Todo struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

/*
	Todo Model对应的数据库表CURD操作函数
*/

func CreateTodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

func GetTodoList() (todoList []*Todo, err error) {
	if err = dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

func GetTodoByID(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err = dao.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateTodo(todo *Todo) (err error) {
	err = dao.DB.Save(&todo).Error
	return
}

func DeleteTodoByID(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
