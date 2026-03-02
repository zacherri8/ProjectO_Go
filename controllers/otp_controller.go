package controllers

import (
	"projecto-backend/services"

	"github.com/gin-gonic/gin"
)

func SendOTP(c *gin.Context) {
	var body struct {
		Email string `json:"email"`
	}
	c.Bind(&body)

	otp, err := services.CreateOTP(body.Email)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to generate OTP"})
		return
	}

	services.SendEmail(body.Email, "Your OTP", "Your OTP is: "+otp)

	c.JSON(200, gin.H{"message": "OTP sent"})
}

func VerifyOTP(c *gin.Context) {
	var body struct {
		Email string `json:"email"`
		OTP   string `json:"otp"`
	}
	c.Bind(&body)

	ok, err := services.VerifyOTP(body.Email, body.OTP)
	if err != nil || !ok {
		c.JSON(400, gin.H{"error": "invalid or expired OTP"})
		return
	}

	services.MarkUserVerified(body.Email)

	c.JSON(200, gin.H{"message": "OTP verified"})
}
