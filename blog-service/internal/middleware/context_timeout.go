package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

func ContextTimeOut(time time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		//创建一个新的ctx
		ctx, cancel := context.WithTimeout(c.Request.Context(), time)
		defer cancel()

		//将新的ctx替换到当前的ctx中
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
