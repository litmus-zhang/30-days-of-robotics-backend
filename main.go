package main

import (
	"30-days-of-robotics-backend/src/database"
	"30-days-of-robotics-backend/src/setup"
	"log"
)

func main() {
	database.Connect()
	database.AutoMigrate()

	err := setup.AppSetup().Run()
	if err != nil {
		log.Printf("Error starting application %v", err.Error())
	}

}
