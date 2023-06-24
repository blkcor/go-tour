package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
	"github.com/go-programming-tour-book/blog-service/pkg/limiter"
	"time"
)

func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		//如果buckets中存在这个key -> 请求已经发生过了
		if bucket, ok := l.GetBucket(key); ok {
			//try take available token
			count := bucket.TakeAvailable(1)
			//no available token -> refuse request
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		} else {
			l.AddBuckets(limiter.LimiterBucketRule{
				Key:          key,
				FillInterval: time.Second,
				Capacity:     10,
				Quantum:      10,
			})
		}
		c.Next()
	}
}
