package routes

import (
	"bwastartup/auth"
	"bwastartup/campaign"
	"bwastartup/config"
	"bwastartup/handler"
	"bwastartup/middleware"
	"bwastartup/transaction"
	"bwastartup/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db, _          = config.ConnectDatabase()

	userRepository = user.NewUserRepository(db)
	userService = user.NewUserService(userRepository)
	userHandler = handler.NewUserHandler(userService, authService)

	authService = auth.NewJWTService()

	campaignRepository = campaign.NewCampaignRepository(db)
	campaignService = campaign.NewCampaignService(campaignRepository)
	campaignHandler =  handler.NewCampaignHandler(campaignService)

	transactionRepository = transaction.NewTransactionRepository(db)
	transactionService = transaction.NewTransactionService(transactionRepository, campaignRepository)
	transactionHandler =  handler.NewTransactionHandler(transactionService)


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
		campaignRoutes.POST("/campaign-images", middleware.AuthorizeJWT(authService, userService), campaignHandler.UploadImage)
	}

	transactionRoutes := r.Group("/api/v1/transactions")

	{
		campaignRoutes.GET("/:id/transactions", middleware.AuthorizeJWT(authService, userService), transactionHandler.GetTransactionsByCampaignId)
		//MELIHAT TRANSAKSI DARI USER YANG DI DAPAT DARI HEADER
		transactionRoutes.GET("/", middleware.AuthorizeJWT(authService, userService), transactionHandler.GetUserTransactionsByUserId)
	}

	r.Static("/gambar", "./images")



	return r
}
