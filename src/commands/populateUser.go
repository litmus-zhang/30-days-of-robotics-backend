package main

import (
	"30-days-of-robotics-backend/src/database"
	"30-days-of-robotics-backend/src/models"
	"github.com/bxcodec/faker/v4"
	"math/rand"
)

func main() {
	database.Connect()
	database.DB.Raw("DELETE from users WHERE track_id IN (1,2,3)").Scan(&models.Task{})

	for i := 0; i < 20; i++ {
		newUser := models.User{
			FirstName: faker.FirstName(),
			LastName:  faker.LastName(),
			Email:     faker.Email(),
			TrackID:   rand.Intn(3+1-1) + 1,
		}
		newUser.SetPassword("123456")
		database.DB.Create(&newUser)
	}
}
