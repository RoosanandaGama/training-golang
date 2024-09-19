package main

import (
	"log"
	"training-golang/session-4-sample-separate-layer/step-4/entity"
	"training-golang/session-4-sample-separate-layer/step-4/handler"
	"training-golang/session-4-sample-separate-layer/step-4/repository/slice"
	"training-golang/session-4-sample-separate-layer/step-4/router"
	"training-golang/session-4-sample-separate-layer/step-4/service"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	var mockUserDBInSlice []entity.User
	userRepo := slice.NewUserRepository(mockUserDBInSlice)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	router.SetupRouter(r, userHandler)

	log.Println("Running server on port 8080")
	r.Run("localhost:8080")
}
