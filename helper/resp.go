package helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespFail(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

func RespOk(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
	})

}
