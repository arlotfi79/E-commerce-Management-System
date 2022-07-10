package Queries

import (
	"API/Communication/DataSignatures"
	"API/Database"
	"log"
)

type ProductQuery struct {
	dbClient *Database.Postgresql
}

func NewProductQuery(dbClient *Database.Postgresql) *ProductQuery {
	return &ProductQuery{dbClient: dbClient}
}

func (productQuery *ProductQuery) GetProductByID(id uint64) (DataSignatures.Product, error) {
	db := productQuery.dbClient.GetDB()

	query, err := db.Prepare(`SELECT * 
									FROM product 
									WHERE product_id = $1`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	row := query.QueryRow(id)

	if err != nil {
		log.Fatal(err)
	}

	var product DataSignatures.Product
	err = row.Scan(&product.Id, &product.Name, &product.Color, &product.Price, &product.Weight, &product.Quantity)

	if err != nil {
		log.Fatal(err)
	}

	return product, nil
}

func (productQuery *ProductQuery) GetProductByName(name string) (DataSignatures.Product, error) {
	db := productQuery.dbClient.GetDB()

	query, err := db.Prepare(`SELECT * 
									FROM product 
									WHERE name = $1`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	row := query.QueryRow(name)

	if err != nil {
		log.Fatal(err)
	}

	var product DataSignatures.Product
	err = row.Scan(&product.Id, &product.Name, &product.Color, &product.Price, &product.Weight, &product.Quantity)

	if err != nil {
		log.Fatal(err)
	}

	return product, nil
}

func (productQuery *ProductQuery) GetProductsByCategoryName(name string) ([]DataSignatures.Product, error) {
	db := productQuery.dbClient.GetDB()

	query, err := db.Prepare(`SELECT p.product_id, p.name, p.color, p.price, p.weight, p.quantity  
									FROM product AS p  
									INNER JOIN product_category AS pc ON p.product_id = pc.product_id  
									INNER JOIN category As c on c.category_id = pc.category_id  
									WHERE c.name = $1`)

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

func (productQuery *ProductQuery) GetProductsByStoreID(id uint64) ([]DataSignatures.Product, error) {
	db := productQuery.dbClient.GetDB()

	query, err := db.Prepare(`SELECT p.product_id, p.name, p.color, p.price, p.weight, p.quantity
									FROM product AS p
									INNER JOIN product_store AS ps ON p.product_id = ps.product_id
									WHERE ps.store_id = $1`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	row, err := query.Query(id)

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

func (productQuery *ProductQuery) GetProductsByOrderID(id uint64) ([]DataSignatures.Product, error) {
	db := productQuery.dbClient.GetDB()
	query, err := db.Prepare(`SELECT p.product_id, p.name, p.color, p.price, p.weight, p.quantity
									FROM product AS p
									INNER JOIN order_product_counter AS op ON p.product_id = op.product_id
									WHERE op.order_id = $1`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	row, err := query.Query(id)

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
