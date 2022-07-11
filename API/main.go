package main

import (

	"API/Database"
	"API/Infrastructure/Handlers"

	middleware "API/Middleware"

	"log"


	"github.com/gin-gonic/gin"
)

func main() {
	var db Database.Postgresql
	err := db.Init()
	defer db.Close()

	if err != nil {
		log.Fatalln(err.Error())
	}

	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())

	// router.POST("/login", userHandler.Login)
	// router.POST("/logout", userHandler.Logout)

	var handlerObj Handlers.AccountHandler

	accountHandler := handlerObj.NewAccountHandler(&db)
	router.POST("/signup", accountHandler.SignUpHandler)
	err = router.Run(":8081")
	if err != nil {
		return
	}
}
