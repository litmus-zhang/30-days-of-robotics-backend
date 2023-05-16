package controller

import (
	"30-days-of-robotics-backend/src/database"
	"30-days-of-robotics-backend/src/middlewares"
	"30-days-of-robotics-backend/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
	userTask := models.UserTask{
		Submitted:  true,
		Submission: data["submission"],
		UserID:     result.ID,
		TrackID:    result.TrackID,
	}
	userTask.SetTaskID(taskID)
	c.JSON(http.StatusAccepted, userTask)
}
