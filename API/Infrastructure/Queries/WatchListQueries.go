package Queries

import (
	"API/Database"
	"log"
)

type WatchListQuery struct {
	dbClient *Database.Postgresql
}

func NewWatchListQuery(dbClient *Database.Postgresql) *WatchListQuery {
	return &WatchListQuery{dbClient: dbClient}
}

func (watchListQuery *WatchListQuery) AddProductToWatchList(accountID uint64, productID uint64) error {
	db := watchListQuery.dbClient.GetDB()

	query, err := db.Prepare(`INSERT INTO WatchList (account_id, product_id)
									VALUES ($1, $2)
									ON CONFLICT 
									    DO NOTHING `)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	if err != nil {
		log.Fatal(err)
	}

	_, err = query.Exec(accountID, productID)

	if err != nil {
		return err
	}

	return nil
}

func (watchListQuery *WatchListQuery) RemoveProductFromWatchList(accountID uint64, productID uint64) error {
	db := watchListQuery.dbClient.GetDB()

	query, err := db.Prepare(`DELETE FROM WatchList
									WHERE account_id = $1 AND product_id = $2`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	if err != nil {
		log.Fatal(err)
	}

	_, err = query.Exec(accountID, productID)

	if err != nil {
		return err
	}

	return nil
}
