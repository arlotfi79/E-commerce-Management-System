package Handlers

import (
	"API/Communication/DataSignatures"
	"API/Database"
	// pwd "API/Infrastructure/PasswordSecurity"
	q "API/Infrastructure/Queries"
	"API/Infrastructure/auth"
	"github.com/gin-gonic/gin"
	// "log"
	"net/http"
)

type ReactiontHandler struct {
	dbClient       *Database.Postgresql
	authInterface  auth.AuthInterface
	tokenInterface auth.TokenInterface
}

func (reactiontHandler *ReactiontHandler) NewReactiontHandler(dbClient *Database.Postgresql, authInt auth.AuthInterface, tokenInt auth.TokenInterface) *ReactiontHandler {
	return &ReactiontHandler{dbClient: dbClient, authInterface: authInt, tokenInterface: tokenInt}
}

func (reactiontHandler *ReactiontHandler) UpVoteHandler(c *gin.Context) {
	var reaction DataSignatures.PostReaction
	reactionq := q.NewReactionQuery(reactiontHandler.dbClient)
	accessInfo, err := reactiontHandler.tokenInterface.ExtractTokenMetadata(c.Request)
	if err == nil {
		if err := c.ShouldBindJSON(&reaction); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"err": "invalid json",
			})
			return
		}
		err = reactionq.UpVoteReaction(reaction.ReviewId, accessInfo.UserId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		return
	} else {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}
}

func (reactiontHandler *ReactiontHandler) DownVoteHandler(c *gin.Context) {
	var reaction DataSignatures.PostReaction
	reactionq := q.NewReactionQuery(reactiontHandler.dbClient)
	accessInfo, err := reactiontHandler.tokenInterface.ExtractTokenMetadata(c.Request)
	if err == nil {
		if err := c.ShouldBindJSON(&reaction); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"err": "invalid json",
			})
			return
		}
		err = reactionq.DownVoteReaction(reaction.ReviewId, accessInfo.UserId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		return
	} else {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}
}