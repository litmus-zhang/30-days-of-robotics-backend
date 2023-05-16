package main

import (
	"30-days-of-robotics-backend/src/database"
	"30-days-of-robotics-backend/src/models"
)

func main() {
	database.Connect()

	for i := 0; i < 3; i++ {
		newTrack := models.Track{
			ID:   uint(i) + 1,
			Name: models.Tracks[i],
		}
		database.DB.Create(&newTrack)
	}
}
