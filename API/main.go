package main

import (
	"API/Database"
	"API/Infrastructure/Handlers"
	"API/Infrastructure/auth"

	middleware "API/Middleware"

	"context"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	var db Database.Postgresql
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	var error error
	redisService, error := auth.NewRedisDB(redisHost, redisPort, redisPassword, context.Background())
	if error != nil {
		log.Fatal(error)
	}

	// authInt := auth.NewAuth()
	tokenInt := auth.NewToken()

	err := db.Init()
	defer db.Close()

	if err != nil {
		log.Fatalln(err.Error())
	}

	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())

	var accHandle Handlers.AccountHandler
	var categHandle Handlers.CategoryHandler
	var prodHandle Handlers.ProductHandler
	var reviewHandle Handlers.ReviewHandler
	var replyHandle Handlers.ReplyHandler
	var tickHandle Handlers.TicketHandler
	var messageHandle Handlers.MessageHandler
	var cartHandle Handlers.CartHandler

	accountHandler := accHandle.NewAccountHandler(&db, redisService.Auth, tokenInt)
	categHandler := categHandle.NewCategoryHandler(&db)
	prodHandler := prodHandle.NewProductHandler(&db, redisService.Auth, tokenInt)
	reviewHandler := reviewHandle.NewReviewHandler(&db, redisService.Auth, tokenInt)
	replyHandler := replyHandle.NewReplyHandler(&db, redisService.Auth, tokenInt)
	tickHandler := tickHandle.NewTicketHandler(&db, redisService.Auth, tokenInt)
	messageHandler := messageHandle.NewTMessageHandler(&db, redisService.Auth, tokenInt)
	cartHandler := cartHandle.NewCartHandler(&db, redisService.Auth, tokenInt)

	router.POST("/signup", accountHandler.SignUpHandler)
	router.POST("/signin", accountHandler.SigninHandler)
	router.GET("/profile", accountHandler.ProfileHandler)

	categoryGroup := router.Group("category")
	{
		categoryGroup.GET("/all", categHandler.GetAllCategoriesHandler)
	}

	productGroup := router.Group("product")
	{
		productGroup.GET("/byCategory", prodHandler.ProductByCategoryNameHandler)
		productGroup.POST("/addNewProduct", prodHandler.AddNewProductHandler)
		productGroup.POST("/addToCategory", prodHandler.AddProductToCategoryHandler)
	}

	cartGroup := router.Group("/cart")
	{
		cartGroup.GET("", cartHandler.GetCartHandler)
		cartGroup.POST("", cartHandler.AddToCartHandler)
	}

	ticketGroup := router.Group("/ticket")
	{
		ticketGroup.GET("", tickHandler.GetTicketsByOrderIDHandler)
		ticketGroup.POST("", tickHandler.PostTicketHandler)
	}

	messageGroup := router.Group("/message")
	{
		messageGroup.GET("", messageHandler.GetMessagesByTicketIDHandler)
		messageGroup.POST("", messageHandler.PostMessageHandler)
	}

	productReviewGroup := router.Group("/productReview")
	{
		productReviewGroup.GET("", reviewHandler.GetReviewsWithVotesByProductIDHandler)
		productReviewGroup.POST("/add", reviewHandler.PostReviewHandler)
		cartGroup.POST("/remove", cartHandler.RemoveFromCartHandler)
	}

	replyReviewGroup := router.Group("/reviewReply")
	{
		replyReviewGroup.GET("", replyHandler.GetRepliesUsingReviewID)
		replyReviewGroup.POST("", replyHandler.PostReply)
	}

	err = router.Run(":8081")
	if err != nil {
		return
	}
}
