package routes

import (
	"projecto-backend/controllers"
	"projecto-backend/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.Engine) {
	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware())
	admin.Use(middleware.AdminAuth())

	admin.GET("/dashboard", controllers.AdminDashboard)
}
