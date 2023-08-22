package router

import (
	"github.com/gin-gonic/gin"
	"morris/im/controller"
	"morris/im/helper"
	"morris/im/middleware"
	"net/http"
	"time"
)

type appConfig struct {
	mode string
}

func InitRouter() *gin.Engine {
	config := initConfig()
	gin.SetMode(config.mode)

	r := gin.New()

	//中间件 记录每次请求时间
	r.Use(middleware.ExecTime())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "restart",
		})
	})

	r.GET("/restart", func(c *gin.Context) {
		time.Sleep(time.Second * 10)
		c.JSON(http.StatusOK, gin.H{
			"msg": "restart success !!!",
		})
	})

	//用户登录注册
	userController := controller.UserController{}
	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)
	r.GET("/search", userController.Search)

	authorized := r.Group("/contact")
	authorized.Use(middleware.Auth())
	{
		//添加好友  好有列表   分组使用中间件
		contactController := controller.ContactController{}
		authorized.POST("/", contactController.AddFriend)
		authorized.GET("/list", contactController.LoadFriend)
	}

	//发起聊天   单个路由 中间件
	chatController := controller.ChatController{}
	r.GET("/chat", middleware.Auth(), chatController.Chat)

	return r
}

func initConfig() appConfig {
	subViper := helper.ViperConfig.Sub("app")
	config := appConfig{
		mode: subViper.GetString("mode"),
	}

	if "" == config.mode {
		config.mode = "debug"
	}

	return config
}
