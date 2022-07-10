package main

import (
	"API/Database"
	middleware "API/Middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	var db Database.Postgresql
	err := db.Init()
	// defer db.Close()

	if err != nil {
		log.Fatalln(err.Error())
	}

	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())

	// router.POST("/login", userHandler.Login)
	// router.POST("/logout", userHandler.Logout)
	var signUpHandler = func(c *gin.Context) {

		return
	}
	router.POST("/signup", signUpHandler)
}
