package main

import (
	"30-days-of-robotics-backend/src/database"
	"30-days-of-robotics-backend/src/models"
	"fmt"
)

func main() {
	database.Connect()
	ans := getAllUserTaskByTrack(2)
	fmt.Println(ans)
}

func getAllUserTaskByTrack(id int) []models.UserTask {
	var UserTask []models.UserTask
	database.DB.Where("track_id = ?", id).Find(&UserTask)
	return UserTask
}
