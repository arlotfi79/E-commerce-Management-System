package Handlers

import (
	"API/Communication/DataSignatures"
	"API/Database"
	q "API/Infrastructure/Queries"
	"API/Infrastructure/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddressHandler struct {
	dbClient       *Database.Postgresql
	authInterface  auth.AuthInterface
	tokenInterface auth.TokenInterface
}

func (addressHandler *AddressHandler) NewAddressHandler(dbClient *Database.Postgresql, authInt auth.AuthInterface, tokenInt auth.TokenInterface) *AddressHandler {
	return &AddressHandler{dbClient: dbClient, authInterface: authInt, tokenInterface: tokenInt}
}

func (addressHandler *AddressHandler) GetAddressByAccountID(c *gin.Context) {
	accessInfo, err := addressHandler.tokenInterface.ExtractTokenMetadata(c.Request)

	if err == nil {
		addressq := q.NewAddressQuery(addressHandler.dbClient)
		addresses, err := addressq.GetAddressesByAccountID(accessInfo.UserId)
		if err != nil {
			c.JSON(http.StatusNotFound, err)
			return
		}

		response := make([]map[string]interface{}, 0, 0)
		for i := 0; i < len(addresses); i++ {
			prodObj := make(map[string]interface{})
			prodObj["id"] = addresses[i].Id
			prodObj["country"] = addresses[i].Country
			prodObj["city"] = addresses[i].City
			prodObj["street"] = addresses[i].Street
			prodObj["plaque"] = addresses[i].Plaque

			response = append(response, prodObj)
		}

		c.JSON(http.StatusOK, response)
	}

}

func (addressHandler *AddressHandler) AddAddressToAccount(c *gin.Context) {
	accessInfo, err := addressHandler.tokenInterface.ExtractTokenMetadata(c.Request)
	var address DataSignatures.PostAddress
	if err == nil {
		if err := c.ShouldBindJSON(&address); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"err": "invalid json",
			})
			return
		}

		addressq := q.NewAddressQuery(addressHandler.dbClient)
		err = addressq.PutAddressUsingAccountID(accessInfo.UserId, &address)
		if err != nil {
			c.JSON(http.StatusNotFound, err)
			return
		}
		return
	} else {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}
}
