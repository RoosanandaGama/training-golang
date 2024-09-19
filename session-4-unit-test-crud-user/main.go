package main

import (
	"log"
	"training-golang/session-4-unit-test-crud-user/entity"
	"training-golang/session-4-unit-test-crud-user/handler"
	"training-golang/session-4-unit-test-crud-user/repository/slice"
	"training-golang/session-4-unit-test-crud-user/router"
	"training-golang/session-4-unit-test-crud-user/service"

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
