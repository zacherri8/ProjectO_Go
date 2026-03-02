package utils

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var limiter = make(map[string][]time.Time)
var lock sync.Mutex

func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		now := time.Now()

		lock.Lock()
		history := limiter[ip]

		// remove old entries
		var fresh []time.Time
		for _, t := range history {
			if now.Sub(t) < time.Minute {
				fresh = append(fresh, t)
			}
		}

		// update map
		limiter[ip] = fresh

		if len(fresh) >= 10 {
			lock.Unlock()
			c.JSON(429, gin.H{"error": "too many requests"})
			c.Abort()
			return
		}

		limiter[ip] = append(limiter[ip], now)
		lock.Unlock()

		c.Next()
	}
}
