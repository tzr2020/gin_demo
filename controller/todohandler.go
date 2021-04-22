package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tzr2020/gin_demo/model"
)

/*
	Todo的处理函数
*/

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodo(c *gin.Context) {
	var todo model.Todo
	// 获取请求包体JSON参数
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 5000,
			"msg":  "failure",
			"err":  err.Error(),
			"data": nil,
		})
		return
	}
	// 数据库创建记录
	if err := model.CreateTodo(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 5000,
			"msg":  "failure",
			"err":  err.Error(),
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func GetTodoList(c *gin.Context) {
	// 数据库查询记录
	todoList, err := model.GetTodoList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 5000,
			"msg":  "failure",
			"err":  err.Error(),
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, todoList)
}

func UpdateTodoByID(c *gin.Context) {
	// 获取请求路径参数
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  4000,
			"msg":   "failure",
			"error": "请求路径参数无效",
			"data":  nil,
		})
		return
	}
	// 数据库查询记录
	todo, err := model.GetTodoByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 5000,
			"msg":  "failure",
			"err":  err.Error(),
			"data": nil,
		})
		return
	}
	// 获取请求包体JSON参数
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  4000,
			"msg":   "failure",
			"error": "请求包体JSON参数无效, err: " + err.Error(),
			"data":  nil,
		})
		return
	}
	// 数据库修改记录
	if err := model.UpdateTodo(todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 5000,
			"msg":  "failure",
			"err":  err.Error(),
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func DeleteTodoByID(c *gin.Context) {
	// 获取请求路径参数
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  4000,
			"msg":   "failure",
			"error": "请求路径参数无效",
			"data":  nil,
		})
		return
	}
	// 数据库删除记录
	if err := model.DeleteTodoByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 5000,
			"msg":  "failure",
			"err":  err.Error(),
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, nil)
}
