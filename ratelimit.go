package ratelimit

import (
	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
)

// RateLimit ratelimit, max qps
func RateLimit(max float64) gin.HandlerFunc {

	lmt := tollbooth.NewLimiter(max, nil)

	return func(c *gin.Context) {
		httpError := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if httpError != nil {
			c.Data(httpError.StatusCode, lmt.GetMessageContentType(), []byte(httpError.Message))
			c.Abort()
		} else {
			c.Next()
		}
	}
}
