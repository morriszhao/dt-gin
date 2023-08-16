package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func ExecTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		execTime := time.Since(startTime)
		log.Println("request time: ", execTime)
	}
}
