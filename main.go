package main

import (
	"github.com/fvbock/endless"
	"log"
	"morris/im/crontab"
	"morris/im/helper"
	"morris/im/models"
	"morris/im/router"
)

func main() {

	//读取配置
	helper.InitConfig()

	//日志启动
	helper.InitLogger()

	//mysql 数据库启动
	models.InitMysql()

	//redis 缓存启动
	helper.InitRedis()

	//定时任务启动
	crontab.InitCronTab()

	//路由启动
	r := router.InitRouter()

	/**
	使用endless 平滑重启  (代码需要从新打包  同名文件)
	默认 endless 会监听一下信号
	syscall.SIGHUP，syscall.SIGUSR1，syscall.SIGUSR2，syscall.SIGINT，syscall.SIGTERM和syscall.SIGTSTP
	接收到 SIGHUP 信号将触发`fork/restart` 实现优雅重启（kill -1 pid会发送SIGHUP信号）
	接收到 syscall.SIGINT或syscall.SIGTERM 信号将触发优雅退出   （服务端收到退出信号不立马退出、等当前处理中的请求全部处理完成在退出）
	*/

	if err := endless.ListenAndServe(":8000", r); err != nil {
		log.Fatal("listen:", err.Error())
	}

	log.Println("server exiting.....")
}
