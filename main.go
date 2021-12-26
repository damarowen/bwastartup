package main

import (
	"bwastartup/handler"
	"bwastartup/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
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
	userHandler := handler.NewUserHandler(userService)

	r := gin.Default()


	userRoutes := r.Group("/api/v1/user", bodySizeMiddleware)
	userRoutes.POST("/register", userHandler.RegisterUser)
	userRoutes.POST("/sessions", userHandler.LoginUser)
	userRoutes.POST("/email_checker", userHandler.IsDuplicateEmail)
	userRoutes.POST("/avatar", userHandler.UploadAvatar)

	err = r.Run(":3000")
	if err != nil {
		log.Fatal(err.Error())
	}
}

//limt upload
func bodySizeMiddleware(c *gin.Context) {
	var w http.ResponseWriter = c.Writer
	c.Request.Body = http.MaxBytesReader(w, c.Request.Body, 1 * 1024 * 1024) // 1 Mb)

	c.Next()
}