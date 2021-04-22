package main

import (
	"log"

	"github.com/tzr2020/gin_demo/dao"
	"github.com/tzr2020/gin_demo/model"
	"github.com/tzr2020/gin_demo/router"
)

func main() {
	// 连接数据库
	err := dao.InitDB()
	if err != nil {
		log.Panicf("connect database failure, err: %v", err)
	}
	defer dao.CloseDB()
	// GORM使用迁移模式
	// Todo Model与数据库表关联
	dao.DB.AutoMigrate(&model.Todo{})

	// 注册路由
	r := router.SetupRouters()

	// 启动HTTP服务
	r.Run()
}
