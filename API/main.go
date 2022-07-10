package main

import (
	"API/Communication/DataSignatures"
	"API/Database"
	q "API/Infrastructure/Queries"

	middleware "API/Middleware"
	"log"
	"net/http"

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
		log.Println("New Request recevied")
		var acc DataSignatures.Account
		if err := c.ShouldBindJSON(&acc); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"invalid_json": "invalid json",
			})
			return
		}
		userq := q.NewUserQuery(&db)
		err := userq.CreateUser(&acc)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		
	}
	router.POST("/signup", signUpHandler)
	err = router.Run("localhost:8081")
	if err != nil {
		return
	}
}
