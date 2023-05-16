package controller

import (
	"30-days-of-robotics-backend/src/database"
	"30-days-of-robotics-backend/src/middlewares"
	"30-days-of-robotics-backend/src/models"
	"30-days-of-robotics-backend/src/services"
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
	}
	task.SetDay(data["day"])
	task.SetTrack(data["track"])

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
		Title:       data["title"],
		Description: data["description"],
	}
	task.SetDay(data["day"])
	task.SetTrack(data["track"])
	database.DB.Model(&task).Where("id = ?", id).Updates(&task)

	c.JSON(http.StatusAccepted, gin.H{"message": "Task updated successfully"})

}

func ViewAllTask(c *gin.Context) {
	id := middlewares.GetUserId(c)

	var result []models.Task
	database.DB.Raw("select u.track_id, t.id, t.day, t.title,t.description "+
		"from tasks t, users u where t.track_id=u.track_id and u.id=?", id).Scan(&result)
	c.JSON(http.StatusAccepted, result)

}

func ViewSingleTask(c *gin.Context) {
	taskID := c.Param("id")
	var task models.Task
	database.DB.Where("id = ?", taskID).First(&task)
	c.JSON(http.StatusOK, task)
}

func SubmitTask(c *gin.Context) {
	var data map[string]string
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid submission")
		return
	}

	userId := middlewares.GetUserId(c)
	type Result struct {
		ID      int
		TrackID int
	}
	var result Result
	database.DB.Raw("select id, track_id from users where id = ?", userId).Scan(&result)
	taskID := c.Param("id")
	check := services.FindTaskByID(taskID)
	if check == false {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No task with such ID"})
		return
	}

	query := database.DB.Where("task_id = ? AND user_id=?", taskID, userId).Find(&models.UserTask{})
	if query.Error != nil || query.RowsAffected != 0 {
		c.JSON(http.StatusBadRequest, "No duplicate submission")
		return
	}

	userTask := models.UserTask{
		Submitted:  true,
		Submission: data["submission"],
		UserID:     result.ID,
		TrackID:    result.TrackID,
	}
	userTask.SetTaskID(taskID)
	database.DB.Create(&userTask)
	c.JSON(http.StatusAccepted, gin.H{"message": "Task submitted successfully"})
}
func GradeTask(c *gin.Context) {
	var data map[string]string
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Grade"})
		return
	}
	taskID := c.Param("id")
	userID := middlewares.GetUserId(c)
	var userTask models.UserTask

	tx := database.DB.Begin()

	database.DB.Where("task_id = ? AND user_id = ?", taskID, userID).Find(&userTask)
	userTask.SetGrade(data["grade"])
	tx.UpdateColumn("grade", userTask.Grade)
	tx.Commit()
	c.JSON(200, userTask)
}
