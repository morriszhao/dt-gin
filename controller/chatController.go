package controller

import (
	"github.com/gin-gonic/gin"
	"morris/im/helper"
	"morris/im/services"
)

type ChatController struct {
}

func (cc *ChatController) Chat(c *gin.Context) {
	err := services.NewChatServices(c).Chat()
	if err != nil {
		helper.RespFail(c, helper.SystemError, err.Error())
		return
	}
}
