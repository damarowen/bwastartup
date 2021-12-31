package main

import (
	"bwastartup/auth"
	"bwastartup/handler"
	"bwastartup/middleware"
	"bwastartup/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {

	dsn := "root:root@tcp(127.0.0.1:3306)/bwagolang?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Connected to DB Success")

	userRepository := user.NewUserRepository(DB)
	userService := user.NewUserService(userRepository)
	authService := auth.NewJWTService()
	userHandler := handler.NewUserHandler(userService, authService)

	r := gin.Default()


	userRoutes := r.Group("/api/v1/user")
	userRoutes.POST("/register", userHandler.RegisterUser)
	userRoutes.POST("/sessions", userHandler.LoginUser)
	userRoutes.POST("/email_checker", userHandler.IsDuplicateEmail)
	userRoutes.POST("/avatar", middleware.AuthorizeJWT(authService),  middleware.BodySizeMiddleware, userHandler.UploadAvatar)

	err = r.Run(":3000")
	if err != nil {
		log.Fatal(err.Error())
	}
}



