package middleware

import (
	"github.com/gin-gonic/gin"
	"morris/im/helper"
	"morris/im/services"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		headerToken := c.Request.Header.Get("token")
		queryToken := c.Request.URL.Query().Get("token")

		if headerToken == "" && queryToken == "" {
			c.Abort()
			helper.RespFail(c, helper.RequestParamsError, "auth fail")
			return
		}

		var token string
		if headerToken != "" {
			token = headerToken
		} else {
			token = queryToken
		}

		userInfo, err := services.NewUserServices(c).GetUserInfoByToken(token)
		if err != nil {
			c.Abort()
			helper.RespFail(c, helper.SystemError, "系统错误")
			return
		}

		if userInfo.Id <= 0 {
			c.Abort()
			helper.RespFail(c, helper.TokenExpired, "token 过期")
			return
		}

		c.Set("user_id", userInfo.Id)
		c.Next()
	}
}
