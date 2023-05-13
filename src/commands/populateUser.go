package main

import (
	"30-days-of-robotics-backend/src/database"
	"30-days-of-robotics-backend/src/models"
)

func main() {
	database.Connect()
	database.DB.Delete(models.User{})

	//for i := 0; i < 5; i++ {
	//	newUser := models.User{
	//		FirstName: faker.FirstName(),
	//		LastName:  faker.LastName(),
	//		Email:     faker.Email(),
	//		Password:  faker.Password(),
	//	}
	//	database.DB.Create(&newUser)
	//}
}
