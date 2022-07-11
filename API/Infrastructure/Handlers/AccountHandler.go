package Handlers

import (
	"API/Communication/DataSignatures"
	"API/Database"
	pwd "API/Infrastructure/PasswordSecurity"
	q "API/Infrastructure/Queries"
	// "log"
	"net/http"
	"API/Infrastructure/auth"
	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	dbClient *Database.Postgresql
	authInterface          auth.AuthInterface
	tokenInterface         auth.TokenInterface
}

func (accountHandler *AccountHandler) NewAccountHandler(dbClient *Database.Postgresql, authInt auth.AuthInterface, tokenInt auth.TokenInterface) *AccountHandler {
	return &AccountHandler{dbClient: dbClient, authInterface: authInt, tokenInterface: tokenInt}
}

func (accountHandler *AccountHandler) SignUpHandler(c *gin.Context) {
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

	err = userq.CreateUser(&acc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

}

func (accountHandler *AccountHandler) SigninHandler(c *gin.Context) {
	var acc DataSignatures.SigninData
	if err := c.ShouldBindJSON(&acc); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"err": "invalid json",
		})
		return
	}
	userq := q.NewUserQuery(accountHandler.dbClient)
	rows, _ := userq.GetUserByUname(acc.UserName)
	if len(rows) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"err": "Username does not exist",
		})
		return
	}

	isMatch := pwd.CheckPasswordHash(acc.Password, rows[0].Password)
	if !isMatch {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err": "Invalid credintials",
		})
		return
	}
	ts, _ := accountHandler.tokenInterface.CreateToken(rows[0].Id)

	userData := make(map[string]interface{})
	userData["access_token"] = ts.AccessToken
	userData["refresh_token"] = ts.RefreshToken
	userData["id"] = rows[0].Id
	userData["first_name"] = rows[0].Name
	userData["last_name"] = rows[0].LastName

	c.JSON(http.StatusOK, userData)


}

