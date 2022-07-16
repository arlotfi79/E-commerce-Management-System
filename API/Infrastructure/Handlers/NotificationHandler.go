package Handlers

import (
	"API/Communication/DataSignatures"
	"API/Database"
	q "API/Infrastructure/Queries"
	"API/Infrastructure/auth"
	"github.com/gin-gonic/gin"
	// "log"
	"net/http"
)

type NotificationHandler struct {
	dbClient       *Database.Postgresql
	authInterface  auth.AuthInterface
	tokenInterface auth.TokenInterface
}

func (notificationHandler *NotificationHandler) NewNotificationHandler(dbClient *Database.Postgresql, authInt auth.AuthInterface, tokenInt auth.TokenInterface) *NotificationHandler {
	return &NotificationHandler{dbClient: dbClient, authInterface: authInt, tokenInterface: tokenInt}
}

func (notificationHandler *NotificationHandler) GetNotificationsHandler(c *gin.Context) {
	var notifs []DataSignatures.Notification
	accessInfo, err := notificationHandler.tokenInterface.ExtractTokenMetadata(c.Request)
	if err == nil {
		notifq := q.NewNotificationQuery(notificationHandler.dbClient)
		notifq.RefreshNotificationsByAccountID(accessInfo.UserId)
		notifs, err = notifq.GetNotificationsByAccountID(accessInfo.UserId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, notifs)
		return
	} else {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}
}