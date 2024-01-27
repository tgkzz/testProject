package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func (h Handler) GinLogger(infoLogger *log.Logger, errLogger *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		duration := time.Since(startTime)

		status := c.Writer.Status()

		infoLogger.Printf("METHOD: %s, PATH: %s, STATUS: %d, DURATION: %v, IP: %s",
			c.Request.Method,
			c.Request.URL.Path,
			status,
			duration,
			c.ClientIP(),
		)
		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				errLogger.Printf("ERROR: %v, PATH: %s, IP: %s",
					e.Err,
					c.Request.URL.Path,
					c.ClientIP(),
				)
			}
		}
	}
}
