package routes

import (
	"projecto-backend/controllers"

	"github.com/gin-gonic/gin"
)

func OTPRoutes(r *gin.Engine) {
	r.POST("/send-otp", controllers.SendOTP)
	r.POST("/verify-otp", controllers.VerifyOTP)
}
