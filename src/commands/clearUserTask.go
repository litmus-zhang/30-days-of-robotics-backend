package main

import (
	"30-days-of-robotics-backend/src/database"
	"30-days-of-robotics-backend/src/models"
)

func main() {
	database.Connect()
	database.DB.Raw("DELETE from user_tasks WHERE track_id IN (1,2,3)").Scan(&models.UserTask{})

}
