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
	var addressHandle Handlers.AddressHandler
	var orderHandle Handlers.OrderHandler
	var watchListHandle Handlers.WatchListHandler
	var NotificationHandle Handlers.NotificationHandler
	var ReactionHandle Handlers.ReactiontHandler

	accountHandler := accHandle.NewAccountHandler(&db, redisService.Auth, tokenInt)
	categHandler := categHandle.NewCategoryHandler(&db, redisService.Auth, tokenInt)
	productHandler := prodHandle.NewProductHandler(&db, redisService.Auth, tokenInt)
	reviewHandler := reviewHandle.NewReviewHandler(&db, redisService.Auth, tokenInt)
	replyHandler := replyHandle.NewReplyHandler(&db, redisService.Auth, tokenInt)
	tickHandler := tickHandle.NewTicketHandler(&db, redisService.Auth, tokenInt)
	messageHandler := messageHandle.NewTMessageHandler(&db, redisService.Auth, tokenInt)
	cartHandler := cartHandle.NewCartHandler(&db, redisService.Auth, tokenInt)
	addressHandler := addressHandle.NewAddressHandler(&db, redisService.Auth, tokenInt)
	orderHandler := orderHandle.NewOrderHandler(&db, redisService.Auth, tokenInt)
	watchListHandler := watchListHandle.NewAddressHandler(&db, redisService.Auth, tokenInt)
	notificationHandler := NotificationHandle.NewNotificationHandler(&db, redisService.Auth, tokenInt)
	reactiontHandler := ReactionHandle.NewReactiontHandler(&db, redisService.Auth, tokenInt)

	router.POST("/signup", accountHandler.SignUpHandler)
	router.POST("/signin", accountHandler.SigninHandler)
	router.GET("/profile", accountHandler.ProfileHandler)

	categoryGroup := router.Group("/category")
	{
		categoryGroup.GET("/all", categHandler.GetAllCategoriesHandler)
		categoryGroup.POST("/addNew", categHandler.AddNewCategoryHandler)
	}

	addressGroup := router.Group("/address")
	{
		addressGroup.POST("/addNew", addressHandler.AddAddressToAccount)
		addressGroup.GET("/getAddresses", addressHandler.GetAddressByAccountID)
	}

	productGroup := router.Group("/product")
	{
		productGroup.GET("/byCategory", productHandler.ProductByCategoryNameHandler)
		productGroup.GET("/watchlist", productHandler.GetWatchListProductsHandler)
		productGroup.POST("/addNewProduct", productHandler.AddNewProductHandler)
		productGroup.POST("/addToCategory", productHandler.AddProductToCategoryHandler)
	}

	cartGroup := router.Group("/cart")
	{
		cartGroup.GET("", cartHandler.GetCartHandler)
		cartGroup.POST("", cartHandler.AddToCartHandler)
	}

	watchListGroup := router.Group("/watchlist")
	{
		watchListGroup.POST("/add", watchListHandler.AddProductToWatchListHandler)
		watchListGroup.POST("/remove", watchListHandler.RemoveProductFromWatchListHandler)
	}

	orderGroup := router.Group("/order")
	{
		orderGroup.GET("/all", orderHandler.GetOrderOfAccountHandler)
		orderGroup.POST("/create", orderHandler.CreateOrderHandler)
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
		productReviewGroup.POST("/upvote", reactiontHandler.UpVoteHandler)
		productReviewGroup.POST("/downvote", reactiontHandler.DownVoteHandler)

		// cartGroup.POST("/remove", cartHandler.RemoveFromCartHandler)
	}

	replyReviewGroup := router.Group("/reviewReply")
	{
		replyReviewGroup.GET("", replyHandler.GetRepliesUsingReviewID)
		replyReviewGroup.POST("", replyHandler.PostReply)
	}
	
	router.GET("/notif", notificationHandler.GetNotificationsHandler)
	
	err = router.Run(":8081")
	if err != nil {
		return
	}
}
