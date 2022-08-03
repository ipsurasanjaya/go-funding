package main

import (
	"log"
	"net/http"
	"strings"

	"crowdfunding/auth"
	"crowdfunding/campaign"
	"crowdfunding/handler"
	"crowdfunding/helper"
	"crowdfunding/payment"
	"crowdfunding/transaction"
	"crowdfunding/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:surasmysql@tcp(127.0.0.1:3306)/crowdfunding_database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	authService := auth.NewService()

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService, authService)

	campaignRepository := campaign.NewRepository(db)
	campaignService := campaign.NewService(campaignRepository)
	campaignHandler := handler.NewCampaignHandler(campaignService)

	transactionRepository := transaction.NewRepository(db)
	paymentService := payment.NewService()
	transactionService := transaction.NewService(transactionRepository, campaignRepository, paymentService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	router := gin.Default()
	router.Static("/images", "./images")

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	api.POST("/campaigns", authMiddleware(authService, userService), campaignHandler.CreateCampaign)
	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaignByID)
	api.PUT("/campaigns/:id", authMiddleware(authService, userService), campaignHandler.UpdateCampaign)
	api.POST("/campaign-images", authMiddleware(authService, userService), campaignHandler.CreateCampaignImage)

	api.GET("/campaigns/:id/transactions", authMiddleware(authService, userService), transactionHandler.GetTransactionByCampaignID)
	api.GET("/transactions", authMiddleware(authService, userService), transactionHandler.GetTransactionByUserID)
	api.POST("/transactions", authMiddleware(authService, userService), transactionHandler.CreateTransaction)
	api.POST("/transactions/notifications", transactionHandler.CreateNotification)
	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			apiResponse := helper.ApiResponse("header token is not contains Bearer", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, apiResponse)
			return
		}

		tokenString := ""
		tokenArray := strings.Split(authHeader, " ")
		if len(tokenArray) == 2 {
			tokenString = tokenArray[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			apiResponse := helper.ApiResponse("Token is not valid", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, apiResponse)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			apiResponse := helper.ApiResponse("Is not ok when converting claims to map claims or token is not valid", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, apiResponse)
			return
		}

		userID := int(claims["user_id"].(float64))
		user, err := userService.GetUserById(userID)
		if err != nil {
			apiResponse := helper.ApiResponse("Error when get user", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, apiResponse)
			return
		}

		c.Set("currentUser", user)
	}
}
