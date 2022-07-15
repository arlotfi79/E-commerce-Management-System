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
	var tickHandle Handlers.TicketHandler
	var messageHandle Handlers.MessageHandler

	accountHandler := accHandle.NewAccountHandler(&db, redisService.Auth, tokenInt)
	categHandler := categHandle.NewCategoryHandler(&db)
	prodHandler := prodHandle.NewProductHandler(&db)
	tickHandler := tickHandle.NewTicketHandler(&db, redisService.Auth, tokenInt)
	messageHandler := messageHandle.NewTMessageHandler(&db, redisService.Auth, tokenInt)

	router.POST("/signup", accountHandler.SignUpHandler)
	router.POST("/signin", accountHandler.SigninHandler)
	router.GET("/profile", accountHandler.ProfileHandler)

	router.GET("/category", categHandler.GetAllCategoriesHandler)
	router.POST("/product", prodHandler.ProductByCategoryNameHandler)

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

	err = router.Run(":8081")
	if err != nil {
		return
	}
}
