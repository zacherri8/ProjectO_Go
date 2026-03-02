package controllers

import (
	"context"

	"projecto-backend/database"
	"projecto-backend/models"

	"github.com/gin-gonic/gin"
)

func AdminDashboard(c *gin.Context) {
	var stats models.AdminStats

	query := `
		SELECT
			(SELECT COUNT(*) FROM users),
			(SELECT COUNT(*) FROM otps)
	`

	err := database.DB.QueryRow(context.Background(), query).Scan(
		&stats.TotalUsers,
		&stats.TotalOtps,
	)

	if err != nil {
		c.JSON(500, gin.H{"error": "database error"})
		return
	}

	c.JSON(200, stats)
}

func GetAllUsers(c *gin.Context) {
	rows, err := database.DB.Query(context.Background(),
		"SELECT id, name, email, phone, branch, year FROM users")
	if err != nil {
		c.JSON(500, gin.H{"error": "database error"})
		return
	}

	var users []models.User
	for rows.Next() {
		var u models.User
		rows.Scan(&u.ID, &u.Name, &u.Email, &u.Phone, &u.Branch, &u.Year)
		users = append(users, u)
	}

	c.JSON(200, users)
}
