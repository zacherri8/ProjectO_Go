package middleware

import (
	"projecto-backend/utils"

	"github.com/gin-gonic/gin"
)

func RateLimitMiddleware() gin.HandlerFunc {
	return utils.RateLimitMiddleware()
}
