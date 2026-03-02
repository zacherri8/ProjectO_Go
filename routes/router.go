package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	AuthRoutes(r)  // /signup, /login, /api/me
	UserRoutes(r)  // /user/profile
	OTPRoutes(r)   // /send-otp, /verify-otp
	AdminRoutes(r) // /admin/dashboard
}
