package main

import (
	"bwastartup/handler"
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
	userHandler := handler.NewUserHandler(userService)

	r := gin.Default()

	userRoutes := r.Group("/api/v1/user")
	userRoutes.POST("/register", userHandler.RegisterUser)
	err = r.Run(":3000")
	if err != nil {
		log.Fatal(err.Error())
	}
}

//func handler(c* gin.Context){
//	dsn := "root:root@tcp(127.0.0.1:3306)/bwagolang?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//
//	if err != nil {
//		log.Fatal(err.Error())
//	}
//	fmt.Println("Connected to DB Success")
//	var users []user.User
//
//	db.Find(&users)
//	c.JSON(http.StatusOK, users)
//	c.JSON(http.StatusOK, "ok")
//
//}
