package Queries

import (
	"API/Communication/DataSignatures"
	"API/Database"
	"log"
)

type AddressQuery struct {
	dbClient *Database.Postgresql
}

func NewAddressQuery(dbClient *Database.Postgresql) *AddressQuery {
	return &AddressQuery{dbClient: dbClient}
}

// -------------------------------- GET ----------------------------------

func (addressQuery *AddressQuery) GetAddressesByAccountID(id string) ([]DataSignatures.Address, error) {
	db := addressQuery.dbClient.GetDB()

	query, err := db.Prepare(`SELECT address_id, country, city, street, plaque
    								FROM address
    								WHERE national_code = $1`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	row, err := query.Query(id)

	if err != nil {
		log.Fatal(err)
	}

	var addresses []DataSignatures.Address
	for row.Next() {
		var address DataSignatures.Address

		err = row.Scan(&address.Id, &address.Country, &address.City, &address.Street, &address.Plaque)

		if err != nil {
			log.Fatal(err)
		}

		addresses = append(addresses, address)
	}

	return addresses, nil
}

// -------------------------------- POST ----------------------------------

func (addressQuery *AddressQuery) PutAddressUsingAccountID(accountID, country, city, street, plaque string) error {
	db := addressQuery.dbClient.GetDB()

	query, err := db.Prepare(`INSERT INTO address (address_id, country, city, street, plaque)
									VALUES ($1, $2, $3, $4, $5)`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	_, err = query.Exec(accountID, country, city, street, plaque)

	if err != nil {
		return err
	}

	return nil
}
