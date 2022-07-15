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

type TicketHandler struct {
	dbClient       *Database.Postgresql
	authInterface  auth.AuthInterface
	tokenInterface auth.TokenInterface
}

func (ticketHandler *TicketHandler) NewTicketHandler(dbClient *Database.Postgresql, authInt auth.AuthInterface, tokenInt auth.TokenInterface) *TicketHandler {
	return &TicketHandler{dbClient: dbClient, authInterface: authInt, tokenInterface: tokenInt}
}

func (ticketHandler *TicketHandler) GetTicketsByOrderIDHandler(c *gin.Context) {
	_, err := ticketHandler.tokenInterface.ExtractTokenMetadata(c.Request)

	if err == nil {
		var order DataSignatures.GetOrder
		if err := c.ShouldBindJSON(&order); err != nil {
			log.Fatal(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"err": "invalid json",
			})
			return
		}

		ticketq := q.NewTicketQuery(ticketHandler.dbClient)
		tickets, err := ticketq.GetTicketsByOrderID(order.Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		response := make([]map[string]interface{}, 0, 0)
		for i := 0; i < len(tickets); i++ {
			prodObj := make(map[string]interface{})
			prodObj["id"] = tickets[i].Id
			prodObj["subject"] = tickets[i].Subject
			prodObj["ticket_date"] = tickets[i].TicketDate

			response = append(response, prodObj)
		}

		c.JSON(http.StatusOK, response)
	}

}

func (ticketHandler *TicketHandler) PostTicketHandler(c *gin.Context) {
	_, err := ticketHandler.tokenInterface.ExtractTokenMetadata(c.Request)

	if err == nil {
		var ticket DataSignatures.PostTicket
		if err := c.ShouldBindJSON(&ticket); err != nil {
			log.Fatal(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"err": "invalid json",
			})
			return
		}

		ticketq := q.NewTicketQuery(ticketHandler.dbClient)

		err := ticketq.PostTicketUsingOrderID(&ticket)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
	}

}
