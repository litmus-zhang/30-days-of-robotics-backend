package main

import (
	"30-days-of-robotics-backend/src/database"
	"30-days-of-robotics-backend/src/models"
	"github.com/bxcodec/faker/v4"
	"math/rand"
)

func main() {
	database.Connect()
	database.DB.Raw("DELETE from tasks WHERE track_id IN (1,2,3)").Scan(&models.Task{})

	for i := 1; i < 31; i++ {
		newTask := models.Task{
			TrackID:     rand.Intn(3-1+1) + 1,
			Day:         i,
			Title:       faker.Word(),
			Description: faker.Username(),
		}
		database.DB.Create(&newTask)
	}

}
