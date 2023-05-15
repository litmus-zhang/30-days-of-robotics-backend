package controller

import (
	"30-days-of-robotics-backend/src/database"
	"30-days-of-robotics-backend/src/models"
	"github.com/bxcodec/faker/v4/pkg/slice"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTask(c *gin.Context) {
	var data map[string]string
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if !slice.Contains(models.Tracks, data["track"]) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Track can only be Programming, Design Or Electronics"})
		return
	}
	query := database.DB.Where("day= ? AND track = ?", data["day"], data["track"]).Find(&models.Task{})

	if query.RowsAffected != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Only single task for a day in a track"})
		return
	}
	task := models.Task{
		Title:       data["title"],
		Description: data["description"],
		Day:         data["day"],
		Track:       data["track"],
	}
	database.DB.Create(&task)

	c.JSON(http.StatusCreated, gin.H{"message": "Task Created Successfully"})
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	query := database.DB.Where("id = ? ", id).First(&models.Task{})
	if query.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No task with such ID"})
		return
	}
	var data map[string]string
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}
	task := models.Task{
		Track:       data["track"],
		Day:         data["day"],
		Title:       data["title"],
		Description: data["description"],
	}
	database.DB.Model(&task).Where("id = ?", id).Updates(&task)

	c.JSON(http.StatusAccepted, gin.H{"message": "Task updated successfully"})

}
