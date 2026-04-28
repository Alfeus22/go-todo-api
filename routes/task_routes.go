package routes

import (
	"GO-1/controllers"
	auth "GO-1/middleware"

	"github.com/gin-gonic/gin"
)

func SetupTaskRouters(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Auth
	protected := r.Group("/Api")
	protected.Use(auth.AuthMiddleware())
	{
		protected.GET("/getTasks", controllers.GetAllTask)
		protected.POST("/tasks", controllers.AddTask)
		protected.DELETE(("/tasks/:id"), controllers.DeleteTask)
		protected.PATCH("/tasks/:id", controllers.UpdateTask)
	}

}
