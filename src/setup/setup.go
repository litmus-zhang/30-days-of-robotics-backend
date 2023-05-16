package setup

import (
	"30-days-of-robotics-backend/src/controller"
	"30-days-of-robotics-backend/src/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func HealthChecker(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "System is up and running"})
	return
}
func AppSetup() *gin.Engine {
	store, _ := redis.NewStore(10, "tcp", "localhost:3000", "", []byte(os.Getenv("JWT_SECRET")))
	store.Options(sessions.Options{MaxAge: 60 * 60 * 24, HttpOnly: true})
	r := gin.Default()
	r.Use(sessions.Sessions("30_DOR", store))
	r.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowAllOrigins:  true,
	}))

	base := r.Group("api/v1")
	r.GET("/health", HealthChecker)

	user := base.Group("users")
	user.POST("/register", controller.Register)
	user.POST("/login", controller.Login)

	authenticatedUser := user.Use(middlewares.IsAuthenticated)
	authenticatedUser.POST("/refresh", controller.RefreshToken)
	authenticatedUser.GET("/user", controller.User)
	authenticatedUser.POST("logout", controller.Logout)
	authenticatedUser.GET("tasks", controller.ViewAllTask)
	authenticatedUser.GET("tasks/:id", controller.ViewSingleTask)
	authenticatedUser.POST("tasks/:id", controller.SubmitTask)
	//authenticatedUser.POST("tasks/:id", controller.GradeTask)

	admin := base.Group("admin")
	adminTask := admin.Group("tasks")
	adminTask.POST("create", controller.CreateTask)
	adminTask.PUT("/:id", controller.UpdateTask)

	return r
}
