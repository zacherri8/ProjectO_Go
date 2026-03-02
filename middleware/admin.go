package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		adminEmail := os.Getenv("ADMIN_EMAIL") // example: admin@kiit.ac.in
		email := c.GetString("email")

		if email != adminEmail {
			c.JSON(http.StatusForbidden, gin.H{"error": "admin only"})
			c.Abort()
			return
		}

		c.Next()
	}
}
