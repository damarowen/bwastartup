package main

import (
	"bwastartup/auth"
	"bwastartup/campaign"
	"bwastartup/handler"
	"bwastartup/middleware"
	"bwastartup/user"
	"fmt"
	"gorm.io/gorm/logger"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:root@tcp(127.0.0.1:3306)/bwagolang?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{ Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Connected to DB Success")

	userRepository := user.NewUserRepository(DB)
	userService := user.NewUserService(userRepository)
	authService := auth.NewJWTService()
	userHandler := handler.NewUserHandler(userService, authService)

	campaignRepository := campaign.NewCampaignRepository(DB)
	campaignService := campaign.NewCampaignService(campaignRepository)
	campaignHandler :=  handler.NewCampaignHandler(campaignService)

	r := gin.Default()


	userRoutes := r.Group("/api/v1/user")
	userRoutes.POST("/register", userHandler.RegisterUser)
	userRoutes.POST("/sessions", userHandler.LoginUser)
	userRoutes.POST("/email_checker", userHandler.IsDuplicateEmail)
	userRoutes.POST("/avatar", middleware.AuthorizeJWT(authService, userService),  middleware.BodySizeMiddleware, userHandler.UploadAvatar)

	campaignRoutes := r.Group("/api/v1/campaign")
	campaignRoutes.GET("", campaignHandler.GetCampaigns)
	campaignRoutes.POST("", middleware.AuthorizeJWT(authService, userService), campaignHandler.CreateCampaign)
	campaignRoutes.GET("/:id", campaignHandler.GetCampaign)
	campaignRoutes.PUT("/:id", middleware.AuthorizeJWT(authService, userService),campaignHandler.UpdateCampaign)

	r.Static("/gambar", "./images")


	err = r.Run(":3000")
	if err != nil {
		log.Fatal(err.Error())
	}
}



