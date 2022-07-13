package main

import (

	"API/Database"
	"API/Infrastructure/Handlers"
	"API/Infrastructure/auth"

	middleware "API/Middleware"

	"log"
	"os"
	"context"
	"github.com/gin-gonic/gin"
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



	accountHandler := accHandle.NewAccountHandler(&db, redisService.Auth, tokenInt)
	categHandler := categHandle.NewCategoryHandler(&db)
	prodHandler := prodHandle.NewProductHandler(&db)

	router.POST("/signup", accountHandler.SignUpHandler)
	router.POST("/signin", accountHandler.SigninHandler)
	router.GET("/category", categHandler.GetAllCategoriesHandler)
	router.POST("/product", prodHandler.ProductByCategoryNameHandler)



	err = router.Run(":8081")
	if err != nil {
		return
	}
}
