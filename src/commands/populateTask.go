package main

import (
	"30-days-of-robotics-backend/src/database"
	"30-days-of-robotics-backend/src/models"
	"github.com/bxcodec/faker/v4"
)

func main() {
	database.Connect()
	database.DB.Delete(&models.Task{}, "track_id IN ?", []int{1, 2, 3})

	for j := 1; j < 4; j++ {
		for i := 1; i < 31; i++ {
			newTask := models.Task{
				TrackID:     j,
				Day:         i,
				Title:       faker.Word(),
				Description: faker.Username(),
			}
			database.DB.Create(&newTask)
		}
	}

}
