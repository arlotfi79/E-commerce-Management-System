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

type ReplyHandler struct {
	dbClient       *Database.Postgresql
	authInterface  auth.AuthInterface
	tokenInterface auth.TokenInterface
}

func (replyHandler *ReplyHandler) NewReplyHandler(dbClient *Database.Postgresql, authInt auth.AuthInterface, tokenInt auth.TokenInterface) *ReplyHandler {
	return &ReplyHandler{dbClient: dbClient, authInterface: authInt, tokenInterface: tokenInt}
}

func (replyHandler *ReplyHandler) GetRepliesUsingReviewID(c *gin.Context) {

	var review DataSignatures.ReviewWithVotes
	if err := c.ShouldBindJSON(&review); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"err": "invalid json",
		})
		return
	}

	replyq := q.NewReplyQuery(replyHandler.dbClient)
	replies, err := replyq.GetRepliesByReviewID(review.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	response := make([]map[string]interface{}, 0, 0)
	for i := 0; i < len(replies); i++ {
		prodObj := make(map[string]interface{})
		prodObj["id"] = replies[i].Id
		prodObj["reply_text"] = replies[i].ReplyText

		response = append(response, prodObj)
	}

	c.JSON(http.StatusOK, response)
}

func (replyHandler *ReplyHandler) PostReply(c *gin.Context) {
	_, err := replyHandler.tokenInterface.ExtractTokenMetadata(c.Request)

	if err == nil {
		var reply DataSignatures.PostReply
		if err := c.ShouldBindJSON(&reply); err != nil {
			log.Fatal(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"err": "invalid json",
			})
			return
		}

		repliq := q.NewReplyQuery(replyHandler.dbClient)

		err := repliq.PostReplyUsingReviewID(&reply)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
	}

}
