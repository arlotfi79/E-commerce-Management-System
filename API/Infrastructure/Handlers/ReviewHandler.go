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

type ReviewHandler struct {
	dbClient       *Database.Postgresql
	authInterface  auth.AuthInterface
	tokenInterface auth.TokenInterface
}

func (reviewHandler *ReviewHandler) NewReviewHandler(dbClient *Database.Postgresql, authInt auth.AuthInterface, tokenInt auth.TokenInterface) *ReviewHandler {
	return &ReviewHandler{dbClient: dbClient, authInterface: authInt, tokenInterface: tokenInt}
}

func (reviewHandler *ReviewHandler) GetReviewsWithVotesByProductIDHandler(c *gin.Context) {
	var product DataSignatures.GetProduct
	if err := c.ShouldBindJSON(&product); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"err": "invalid json",
		})
		return
	}

	reviewq := q.NewReviewQuery(reviewHandler.dbClient)
	reviews, err := reviewq.GetReviewsWithVotesByProductID(product.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	response := make([]map[string]interface{}, 0, 0)
	for i := 0; i < len(reviews); i++ {
		prodObj := make(map[string]interface{})
		prodObj["id"] = reviews[i].Id
		prodObj["description"] = reviews[i].Description
		prodObj["rating"] = reviews[i].Rating
		prodObj["up_vote_count"] = reviews[i].UpVoteCount
		prodObj["down_vote_count"] = reviews[i].DownVoteCount
		prodObj["rating_avg"] = reviews[i].RatingAVG

		response = append(response, prodObj)
	}

	c.JSON(http.StatusOK, response)
}

func (reviewHandler *ReviewHandler) PostReviewHandler(c *gin.Context) {
	_, err := reviewHandler.tokenInterface.ExtractTokenMetadata(c.Request)

	if err == nil {
		var review DataSignatures.PostReview
		if err := c.ShouldBindJSON(&review); err != nil {
			log.Fatal(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"err": "invalid json",
			})
			return
		}

		reviewq := q.NewReviewQuery(reviewHandler.dbClient)

		err := reviewq.PostReviewUsingProductID(&review)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
	}
}
