package main

import (
	"30-days-of-robotics/Setup"
	"30-days-of-robotics/src/database"
)

func main() {
	Setup.AppSetup()
	database.Connect()
	database.AutoMigrate()
}
