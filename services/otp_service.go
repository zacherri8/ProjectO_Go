package services

import (
	"context"
	"projecto-backend/database"
	"projecto-backend/models"
	"projecto-backend/utils"
	"time"
)

func CreateOTP(email string) (string, error) {
	otp := utils.GenerateOTP()

	_, err := database.DB.Exec(
		context.Background(),
		`INSERT INTO otps (email, otp, expires_at)
		 VALUES ($1, $2, $3)
		 ON CONFLICT (email) 
		 DO UPDATE SET otp=$2, expires_at=$3`,
		email,
		otp,
		time.Now().Add(5*time.Minute),
	)

	return otp, err
}

func VerifyOTP(email string, otp string) bool {
	var stored models.OTP

	err := database.DB.QueryRow(
		context.Background(),
		"SELECT email, otp, expires_at FROM otps WHERE email=$1",
		email,
	).Scan(&stored.Email, &stored.OTP, &stored.ExpiresAt)

	if err != nil {
		return false
	}

	if stored.OTP != otp {
		return false
	}

	return time.Now().Before(stored.ExpiresAt)
}
