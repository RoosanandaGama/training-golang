package router

import (
	"training-golang/session-2-latihan-crud-user-gin/handler"
	"training-golang/session-2-latihan-crud-user-gin/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	userPublicEndpoint := r.Group("/user")
	userPublicEndpoint.GET("/", handler.GetAllUsers)
	userPublicEndpoint.GET("/:id", handler.GetUser)

	userPrivateEndpoint := r.Group("/user")
	userPrivateEndpoint.Use(middleware.AuthMiddleware())
	{
		userPrivateEndpoint.POST("/", handler.CreatedUser)
		userPrivateEndpoint.PUT("/:id", handler.UpdateUser)
		userPrivateEndpoint.DELETE("/:id", handler.DeleteUser)
	}
}
