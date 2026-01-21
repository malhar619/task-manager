package controllers

import (
	"fmt"
	"net/http"
	"task-manager/config"
	"task-manager/models"
	"task-manager/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// 1. Bind JSON from Postman
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return
	}

	// 2. Find User in DB
	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// 3. Compare Passwords (CRITICAL STEP)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		fmt.Println("Bcrypt Error:", err) // This will show in your terminal
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// 4. Generate Token
	token, _ := utils.GenerateToken(user.ID, user.Role)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Register(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	input.Password = string(hashed)

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists or database error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
