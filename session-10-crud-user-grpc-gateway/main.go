package main

import (
	"context"
	"log"
	"net"
	grpchandler "training-golang/session-10-crud-user-grpc-gateway/handler/grpc"
	"training-golang/session-10-crud-user-grpc-gateway/middleware"
	pb "training-golang/session-10-crud-user-grpc-gateway/proto/user_service/v1"
	postgresgorm "training-golang/session-10-crud-user-grpc-gateway/repository/postgres_gorm"
	"training-golang/session-10-crud-user-grpc-gateway/service"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	dsn := "postgres://postgres:P4ssw0rd@192.168.26.50:5432/latihangolang"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})

	if err != nil {
		log.Fatalln(err)
	}

	//setup service
	userRepo := postgresgorm.NewUserRepository(gormDB)
	userService := service.NewUserService(userRepo)
	userHandler := grpchandler.NewUserHandler(userService)

	//run the grpc server
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(middleware.UnaryAuthInterceptor()))
	pb.RegisterUserServiceServer(grpcServer, userHandler)

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		log.Println("Server dah jalan kok bos di port 50051")
		grpcServer.Serve(lis)
	}()

	//run the grpc-gateway
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	if err := pb.RegisterUserServiceHandler(context.Background(), gwmux, conn); err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	//run gin server

	ginServer := gin.Default()

	ginServer.Group("/v1/*{grpc_gateway}").Any("", gin.WrapH(gwmux))

	log.Println("Server dah jalan bos di port 8080")

	ginServer.Run("localhost:8080")

}
