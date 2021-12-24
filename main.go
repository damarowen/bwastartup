package main

import (
	"bwastartup/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {

	//dsn := "root:root@tcp(127.0.0.1:3306)/bwagolang?charset=utf8mb4&parseTime=True&loc=Local"
	//_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//fmt.Println("Connected to DB Success")

	router := gin.Default()
	router.GET("/", handler)
	err := router.Run(":3000")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func handler(c* gin.Context){
	dsn := "root:root@tcp(127.0.0.1:3306)/bwagolang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Connected to DB Success")
	var users []user.User

	db.Find(&users)
	c.JSON(http.StatusOK, users)
}
