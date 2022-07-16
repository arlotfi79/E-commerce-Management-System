package Handlers

import (
	"API/Communication/DataSignatures"
	"API/Database"
	q "API/Infrastructure/Queries"
	"API/Infrastructure/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderHandler struct {
	dbClient       *Database.Postgresql
	authInterface  auth.AuthInterface
	tokenInterface auth.TokenInterface
}

func (orderHandler *OrderHandler) NewOrderHandler(dbClient *Database.Postgresql, authInt auth.AuthInterface, tokenInt auth.TokenInterface) *OrderHandler {
	return &OrderHandler{dbClient: dbClient, authInterface: authInt, tokenInterface: tokenInt}
}

func (orderHandler *OrderHandler) GetOrderOfAccountHandler(c *gin.Context) {
	accessInfo, err := orderHandler.tokenInterface.ExtractTokenMetadata(c.Request)

	if err == nil {
		orderq := q.NewOrderQuery(orderHandler.dbClient)
		orders, err := orderq.GetOrdersByAccountID(accessInfo.UserId)
		if err != nil {
			c.JSON(http.StatusNotFound, err)
			return
		}

		response := make([]map[string]interface{}, 0, 0)
		for i := 0; i < len(orders); i++ {
			prodObj := make(map[string]interface{})
			prodObj["id"] = orders[i].Id
			prodObj["description"] = orders[i].Description
			prodObj["deliveryMethod"] = orders[i].DeliveryMethod
			prodObj["orderDate"] = orders[i].OrderDate
			prodObj["address"] = orders[i].Address

			response = append(response, prodObj)
		}

		c.JSON(http.StatusOK, response)
	}
}

func (orderHandler *OrderHandler) CreateOrderHandler(c *gin.Context) {
	accessInfo, err := orderHandler.tokenInterface.ExtractTokenMetadata(c.Request)
	var order DataSignatures.PostOrder

	if err == nil {
		if err := c.ShouldBindJSON(&order); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"err": "invalid json",
			})
			return
		}

		orderq := q.NewOrderQuery(orderHandler.dbClient)
		err = orderq.CreateOrder(accessInfo.UserId, order)
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
