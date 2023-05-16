package services

import (
	"30-days-of-robotics-backend/src/database"
	"30-days-of-robotics-backend/src/models"
)

func FindTaskByID(id string) bool {
	query := database.DB.Where("id = ? ", id).First(&models.Task{})
	if query.Error != nil {
		return false
	}
	return true
}
