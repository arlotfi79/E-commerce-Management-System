package Queries

import (
	"API/Database"
	"log"
)

type CartQuery struct {
	dbClient *Database.Postgresql
}

func NewCartQuery(dbClient *Database.Postgresql) *CartQuery {
	return &CartQuery{dbClient: dbClient}
}

// -------------------------------- POST ----------------------------------

func (cartListQuery *CartQuery) AddProductToCartList(accountID uint64, productID uint64, productCount uint64) error {
	db := cartListQuery.dbClient.GetDB()

	query, err := db.Prepare(`INSERT INTO Cart (account_id, product_id, product_count)
									VALUES ($1, $2, $3)
									ON CONFLICT 
									    DO NOTHING `)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	if err != nil {
		log.Fatal(err)
	}

	_, err = query.Exec(accountID, productID, productCount)

	if err != nil {
		return err
	}

	return nil
}

func (cartListQuery *CartQuery) RemoveProductFromCartList(accountID uint64, productID uint64) error {
	db := cartListQuery.dbClient.GetDB()

	query, err := db.Prepare(`DELETE FROM Cart
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

func (cartListQuery *CartQuery) IncreaseCountOfProductInCartList(accountID uint64, productID uint64) error {
	db := cartListQuery.dbClient.GetDB()

	query, err := db.Prepare(`UPDATE Cart
									SET product_count = product_count + 1
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

func (cartListQuery *CartQuery) DecreaseCountOfProductInCartList(accountID uint64, productID uint64) error {
	db := cartListQuery.dbClient.GetDB()

	query, err := db.Prepare(`UPDATE Cart
									SET product_count = product_count - 1
									WHERE account_id = $1 AND product_id = $2 AND product_count > 0`)

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
