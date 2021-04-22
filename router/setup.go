package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tzr2020/gin_demo/controller"
)

func SetupRouters() (r *gin.Engine) {
	// 使用Gin默认路由器
	r = gin.Default()

	// 静态资源处理器
	r.Static("/static", "static")

	// 加载模板文件
	r.LoadHTMLFiles("template/index.html")
	// 渲染模板
	r.GET("/", controller.IndexHandler)

	// api v1
	v1Group := r.Group("v1")
	{
		// Todo的路由注册

		// 添加待办事项
		v1Group.POST("/todo", controller.CreateTodo)
		// 查看所有待办事项
		v1Group.GET("todo", controller.GetTodoList)
		// 修改一个待办事项
		v1Group.PUT("todo/:id", controller.UpdateTodoByID)
		// 删除一个待办事项
		v1Group.DELETE("todo/:id", controller.DeleteTodoByID)
	}

	return
}
