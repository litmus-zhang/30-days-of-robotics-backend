package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "Registration successful",
	})
}

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
	})
}
