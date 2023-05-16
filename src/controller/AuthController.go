package controller

import (
	"30-days-of-robotics-backend/src/database"
	"30-days-of-robotics-backend/src/middlewares"
	"30-days-of-robotics-backend/src/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"time"
)

func Register(c *gin.Context) {
	var data map[string]string
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if data["confirm_password"] != data["password"] {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Password Mismatch"})
		return
	}
	type Result struct {
		Email string
	}
	var result Result
	query := database.DB.Raw("SELECT email FROM users WHERE email = ? ", data["email"]).Scan(&result)
	if query.RowsAffected != 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Email already in use"})
		return
	}
	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}
	user.SetPassword(data["password"])
	user.SetTrack(data["track"])
	database.DB.Create(&user)
	c.JSON(http.StatusCreated, gin.H{"message": "User registration successful"})
	return
}

func Login(c *gin.Context) {
	var data map[string]string
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid credential"})
		return
	}
	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid credential"})
		return
	}

	err := user.ComparePassword(data["password"])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid credential"})
		return
	}
	payload := jwt.StandardClaims{
		Subject:   strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, er := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid credential"})
		return
	}
	sessionToken := user.ID
	session := sessions.Default(c)
	session.Set("userID", sessionToken)
	session.Set("token", tokenString)
	err = session.Save()
	if err != nil {
		panic("Error saving session token")
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "Login Successful"})
	return
}
func User(c *gin.Context) {
	id := middlewares.GetUserId(c)

	var user models.User
	database.DB.Where("id= ?", id).First(&user)
	c.JSON(200, user)
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)

	session.Delete("token")
	session.Delete("userID")
	err := session.Save()
	if err != nil {
		panic("Error saving session")
	}
	c.JSON(200, gin.H{"message": "Logout successful"})
	return
}
func RefreshToken(c *gin.Context) {
	token, _ := c.Cookie("jwt")
	c.JSON(200, gin.H{"token": token})

}
