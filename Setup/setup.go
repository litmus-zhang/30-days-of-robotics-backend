package Setup

import (
	"30-days-of-robotics/src/controller"
	"github.com/gin-gonic/gin"
	"log"
)

func AppSetup() {
	r := gin.Default()
	base := r.Group("api/v1")
	user := base.Group("users")
	user.POST("/login", controller.Login)
	user.POST("/register", controller.Register)

	err := r.Run()
	if err != nil {
		log.Printf("Error starting application %v", err.Error())
	}

}
