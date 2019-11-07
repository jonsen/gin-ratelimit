# gin-ratelimit
gin-ratelimit, a ratelimit middleware for gin.


example:
```go
package main

import (
	"github.com/gin-gonic/gin"
	ratelimit "github.com/jonsen/gin-ratelimit"
)

func main() {
	r := gin.New()

	r.Use(ratelimit.RateLimit(20))

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, world!")
	})

	r.Run(":12345")
}
```
