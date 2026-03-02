package middleware

import (
	"projecto-backend/utils"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return utils.RequestLogger()
}
