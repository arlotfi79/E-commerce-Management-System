package Handlers

import (
	"API/Communication/DataSignatures"
	"API/Database"
	q "API/Infrastructure/Queries"
	"API/Infrastructure/auth"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type WatchListHandler struct {
	dbClient       *Database.Postgresql
	authInterface  auth.AuthInterface
	tokenInterface auth.TokenInterface
}

func (watchListHandler *WatchListHandler) NewAddressHandler(dbClient *Database.Postgresql, authInt auth.AuthInterface, tokenInt auth.TokenInterface) *WatchListHandler {
	return &WatchListHandler{dbClient: dbClient, authInterface: authInt, tokenInterface: tokenInt}
}

func (watchListHandler *WatchListHandler) AddProductToWatchListHandler(c *gin.Context) {
	accessInfo, err := watchListHandler.tokenInterface.ExtractTokenMetadata(c.Request)

	if err == nil {
		var product DataSignatures.PostProductID
		if err := c.ShouldBindJSON(&product); err != nil {
			log.Fatal(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"err": "invalid json",
			})
			return
		}

		watchListq := q.NewWatchListQuery(watchListHandler.dbClient)

		err := watchListq.AddProductToWatchList(accessInfo.UserId, product.Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
	}
}

func (watchListHandler *WatchListHandler) RemoveProductFromWatchListHandler(c *gin.Context) {
	accessInfo, err := watchListHandler.tokenInterface.ExtractTokenMetadata(c.Request)

	if err == nil {
		var product DataSignatures.PostProductID
		if err := c.ShouldBindJSON(&product); err != nil {
			log.Fatal(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"err": "invalid json",
			})
			return
		}

		watchListq := q.NewWatchListQuery(watchListHandler.dbClient)

		err := watchListq.RemoveProductFromWatchList(accessInfo.UserId, product.Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
	}
}
