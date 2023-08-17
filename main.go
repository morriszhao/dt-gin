package main

import (
	"morris/im/helper"
	"morris/im/models"
	"morris/im/router"
	"net/http"
)

func main() {

	//读取配置
	helper.InitConfig()

	//mysql 数据库启动
	models.InitMysql()

	//redis 缓存启动
	helper.InitRedis()

	//路由启动
	router := router.InitRouter()

	//启动服务
	httpserver := http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	httpserver.ListenAndServe()
}
