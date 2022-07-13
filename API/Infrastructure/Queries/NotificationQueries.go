package Queries

import (
	"API/Communication/DataSignatures"
	"API/Database"
	"log"
)

type NotificationQuery struct {
	dbClient *Database.Postgresql
}

func NewNotificationQuery(dbClient *Database.Postgresql) *NotificationQuery {
	return &NotificationQuery{dbClient: dbClient}
}

func (notificationQuery *NotificationQuery) GetNotificationsByAccountID(accountID uint64) ([]DataSignatures.Notification, error) {
	db := notificationQuery.dbClient.GetDB()

	query, err := db.Prepare(`SELECT *
									FROM Notification
									WHERE account_id = $1`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	row, err := query.Query(accountID)

	if err != nil {
		log.Fatal(err)
	}

	var notifications []DataSignatures.Notification
	for row.Next() {
		var notification DataSignatures.Notification

		err = row.Scan()

		if err != nil {
			log.Fatal(err)
		}

		notifications = append(notifications, notification)
	}

	return notifications, nil
}

func (notificationQuery *NotificationQuery) RefreshNotificationsByAccountID(accountID uint64) error {
	db := notificationQuery.dbClient.GetDB()

	query, err := db.Prepare(`CALL RefreshNotificationsAndClearAvailableWatchListProducts($1)`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	_, err = query.Exec(accountID)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
