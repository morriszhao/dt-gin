package services

import "github.com/gin-gonic/gin"

type BaseServices struct {
}

// GetUserId 获取userId
func GetUserId(c *gin.Context) int {
	value, exists := c.Get("user_id")
	if !exists {
		return 0
	}

	userId := value.(int64)
	return int(userId)
}
