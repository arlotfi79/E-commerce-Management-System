package Queries

import (
	"DBProject/API/Communication/DataSignatures"
	"DBProject/API/Database"
	"log"
)

type ProductQuery struct {
	dbClient *Database.Postgresql
}

func NewProductQuery(dbClient *Database.Postgresql) *ProductQuery {
	return &ProductQuery{dbClient: dbClient}
}

func (productQuery *ProductQuery) getProductsByCategory(name string) ([]DataSignatures.Product, error) {
	db := productQuery.dbClient.GetDB()

	query, err := db.Prepare("SELECT *" +
		"FROM product AS p" +
		"INNER JOIN product_category AS pc ON p.product_id = pc.product_id" +
		"INNER JOIN category As c on c.category_id = pc.category_id" +
		"WHERE c.name = ?")

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	row, err := query.Query(name)

	if err != nil {
		log.Fatal(err)
	}

	var products []DataSignatures.Product
	for row.Next() {
		var product DataSignatures.Product
		err = row.Scan(&product.Id, &product.Name, &product.Color, &product.Price, &product.Weight, &product.Quantity)

		if err != nil {
			log.Fatal(err)
		}

		products = append(products, product)
	}

	return products, nil
}
