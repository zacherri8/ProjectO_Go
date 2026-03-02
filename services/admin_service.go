package services

import (
	"context"
	"projecto-backend/database"
	"projecto-backend/models"
)

func GetAdminStats() (models.AdminStats, error) {
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

	return stats, err
}
