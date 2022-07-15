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

type MessageHandler struct {
	dbClient       *Database.Postgresql
	authInterface  auth.AuthInterface
	tokenInterface auth.TokenInterface
}

func (messageHandler *MessageHandler) NewTMessageHandler(dbClient *Database.Postgresql, authInt auth.AuthInterface, tokenInt auth.TokenInterface) *MessageHandler {
	return &MessageHandler{dbClient: dbClient, authInterface: authInt, tokenInterface: tokenInt}
}

func (messageHandler *MessageHandler) GetMessagesByTicketIDHandler(c *gin.Context) {
	_, err := messageHandler.tokenInterface.ExtractTokenMetadata(c.Request)

	if err == nil {
		var ticket DataSignatures.Ticket
		if err := c.ShouldBindJSON(&ticket); err != nil {
			log.Fatal(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"err": "invalid json",
			})
			return
		}

		messageq := q.NewMessageQuery(messageHandler.dbClient)
		messages, err := messageq.GetMessagesByTicketID(ticket.Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		response := make([]map[string]interface{}, 0, 0)
		for i := 0; i < len(messages); i++ {
			prodObj := make(map[string]interface{})
			prodObj["id"] = messages[i].Id
			prodObj["message_text"] = messages[i].MessageText
			prodObj["message_date"] = messages[i].MessageDate

			response = append(response, prodObj)
		}

		c.JSON(http.StatusOK, response)
	}

}

func (messageHandler *MessageHandler) PostMessageHandler(c *gin.Context) {
	_, err := messageHandler.tokenInterface.ExtractTokenMetadata(c.Request)

	if err == nil {
		var message DataSignatures.PostMessage
		if err := c.ShouldBindJSON(&message); err != nil {
			log.Fatal(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"err": "invalid json",
			})
			return
		}

		ticketq := q.NewMessageQuery(messageHandler.dbClient)

		err := ticketq.PostMessageUsingTicketID(&message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
	}

}
