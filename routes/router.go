package routes

import (
	"bwastartup/auth"
	"bwastartup/campaign"
	"bwastartup/config"
	"bwastartup/handler"
	"bwastartup/middleware"
	"bwastartup/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db, _          = config.ConnectDatabase()
	userRepository = user.NewUserRepository(db)
	userService = user.NewUserService(userRepository)
	authService = auth.NewJWTService()
	userHandler = handler.NewUserHandler(userService, authService)

	campaignRepository = campaign.NewCampaignRepository(db)
	campaignService = campaign.NewCampaignService(campaignRepository)
	campaignHandler =  handler.NewCampaignHandler(campaignService)

)

//SetupRouter ... Configure routes
func SetupRouter(db *gorm.DB) *gin.Engine {

	r := gin.Default()

	userRoutes := r.Group("/api/v1/user")
	{
		userRoutes.POST("/register", userHandler.RegisterUser)
		userRoutes.POST("/sessions", userHandler.LoginUser)
		userRoutes.POST("/email_checker", userHandler.IsDuplicateEmail)
		userRoutes.POST("/avatar", middleware.AuthorizeJWT(authService, userService),  middleware.BodySizeMiddleware, userHandler.UploadAvatar)
	}

	campaignRoutes := r.Group("/api/v1/campaign")
	{
		campaignRoutes.GET("", campaignHandler.GetCampaigns)
		campaignRoutes.POST("", middleware.AuthorizeJWT(authService, userService), campaignHandler.CreateCampaign)
		campaignRoutes.GET("/:id", campaignHandler.GetCampaign)
		campaignRoutes.PUT("/:id", middleware.AuthorizeJWT(authService, userService),campaignHandler.UpdateCampaign)
		campaignRoutes.POST("/campaign-images", middleware.AuthorizeJWT(authService, userService),campaignHandler.UploadImage)
	}

	r.Static("/gambar", "./images")



	return r
}
