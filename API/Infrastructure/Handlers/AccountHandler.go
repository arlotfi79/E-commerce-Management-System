package Handlers

import (
	"API/Communication/DataSignatures"
	"API/Database"
	pwd "API/Infrastructure/PasswordSecurity"
	q "API/Infrastructure/Queries"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	dbClient *Database.Postgresql
}

func (accountHandler *AccountHandler) NewAccountHandler(dbClient *Database.Postgresql) *AccountHandler {
	return &AccountHandler{dbClient: dbClient}
}

func (accountHandler *AccountHandler) SignUpHandler(c *gin.Context) {
	log.Println("New Request recevied")
	var acc DataSignatures.PostAccount
	if err := c.ShouldBindJSON(&acc); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"err": "invalid json",
		})
		return
	}
	userq := q.NewUserQuery(accountHandler.dbClient)

	//Check if user is duplicate
	accounts, err := userq.GetUserByUname(acc.UserName)
	if len(accounts) != 0 {
		c.JSON(http.StatusForbidden, gin.H{
			"err": "Duplicate username",
		})
		return
	}

	accounts, err = userq.GetUserByEmail(acc.Email)
	if len(accounts) != 0 {
		c.JSON(http.StatusForbidden, gin.H{
			"err": "Duplicate email",
		})
		return
	}

	accounts, err = userq.GetUserByPhoneNumber(acc.PhoneNumber)
	if len(accounts) != 0 {
		c.JSON(http.StatusForbidden, gin.H{
			"err": "Duplicate phoneNumber",
		})
		return
	}

	//Hashing password
	hashedPass, _ := pwd.Encrypt(acc.Password)
	acc.Password = string(hashedPass)
	err = userq.CreateUser(&acc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

}
