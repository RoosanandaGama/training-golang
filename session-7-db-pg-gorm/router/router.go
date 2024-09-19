package router

import (
	"training-golang/session-7-db-pg-gorm/handler"
	"training-golang/session-7-db-pg-gorm/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, userHandler handler.IUserHandler) {
	userPublicEndpoint := r.Group("/users")

	userPublicEndpoint.GET("/:id", userHandler.GetUser)
	userPublicEndpoint.GET("/", userHandler.GetAllUsers)
	userPublicEndpoint.GET("", userHandler.GetAllUsers)

	userPrivateEndpoint := r.Group("/users")
	userPrivateEndpoint.Use(middleware.AuthMiddleware())
	userPrivateEndpoint.POST("/", userHandler.CreateUser)
	userPrivateEndpoint.POST("", userHandler.CreateUser)
	userPrivateEndpoint.PUT("/:id", userHandler.UpdateUser)
	userPrivateEndpoint.DELETE("/:id", userHandler.DeleteUser)
}
