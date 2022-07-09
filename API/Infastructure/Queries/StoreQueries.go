package Queries

import (
	"API/Communication/DataSignatures"
	"API/Database"
	"log"
)

type StoreQuery struct {
	dbClient *Database.Postgresql
}

func NewStoreQuery(dbClient *Database.Postgresql) *StoreQuery {
	return &StoreQuery{dbClient: dbClient}
}

func (storeQuery *StoreQuery) GetStoreByID(id uint64) (DataSignatures.Store, error) {
	db := storeQuery.dbClient.GetDB()

	query, err := db.Prepare(`SELECT * 
									FROM store 
									WHERE store_id = $1`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	row := query.QueryRow(id)

	if err != nil {
		log.Fatal(err)
	}

	var store DataSignatures.Store
	err = row.Scan(&store.Id, &store.Name, &store.PhoneNumber)

	if err != nil {
		log.Fatal(err)
	}

	return store, nil
}
