package routes

import (
	"task-manager/controllers"
	"task-manager/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/tasks", controllers.CreateTask)       // [cite: 11]
		protected.GET("/tasks", controllers.GetTasks)          // [cite: 11]
		protected.GET("/tasks/:id", controllers.GetTaskByID)   // [cite: 12]
		protected.DELETE("/tasks/:id", controllers.DeleteTask) // [cite: 14]
	}
}
