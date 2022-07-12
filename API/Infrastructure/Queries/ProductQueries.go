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

// -------------------------------- GET ----------------------------------

func (productQuery *ProductQuery) GetProductByID(id uint64) (DataSignatures.GetProduct, error) {
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

	var product DataSignatures.GetProduct
	err = row.Scan(&product.Id, &product.Name, &product.Color, &product.Price, &product.Weight, &product.Quantity)

	if err != nil {
		log.Fatal(err)
	}

	return product, nil
}

func (productQuery *ProductQuery) GetProductByName(name string) (DataSignatures.GetProduct, error) {
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

	var product DataSignatures.GetProduct
	err = row.Scan(&product.Id, &product.Name, &product.Color, &product.Price, &product.Weight, &product.Quantity)

	if err != nil {
		log.Fatal(err)
	}

	return product, nil
}

func (productQuery *ProductQuery) GetProductsByCategoryName(name string) ([]DataSignatures.GetProduct, error) {
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

	var products []DataSignatures.GetProduct
	for row.Next() {
		var product DataSignatures.GetProduct
		err = row.Scan(&product.Id, &product.Name, &product.Color, &product.Price, &product.Weight, &product.Quantity)

		if err != nil {
			log.Fatal(err)
		}

		products = append(products, product)
	}

	return products, nil
}

func (productQuery *ProductQuery) GetProductsByStoreID(id uint64) ([]DataSignatures.GetProduct, error) {
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

	var products []DataSignatures.GetProduct
	for row.Next() {
		var product DataSignatures.GetProduct
		err = row.Scan(&product.Id, &product.Name, &product.Color, &product.Price, &product.Weight, &product.Quantity)

		if err != nil {
			log.Fatal(err)
		}

		products = append(products, product)
	}

	return products, nil
}

func (productQuery *ProductQuery) GetProductsByOrderID(id uint64) ([]DataSignatures.GetProduct, error) {
	db := productQuery.dbClient.GetDB()
	query, err := db.Prepare(`SELECT p.product_id, p.name, p.color, p.price, p.weight, p.quantity
									FROM product AS p
									INNER JOIN order_product AS op ON p.product_id = op.product_id
									WHERE op.order_id = $1`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	row, err := query.Query(id)

	if err != nil {
		log.Fatal(err)
	}

	var products []DataSignatures.GetProduct
	for row.Next() {
		var product DataSignatures.GetProduct
		err = row.Scan(&product.Id, &product.Name, &product.Color, &product.Price, &product.Weight, &product.Quantity)

		if err != nil {
			log.Fatal(err)
		}

		products = append(products, product)
	}

	return products, nil
}

func (productQuery *ProductQuery) GetProductsOfWatchList(accountID uint64, productID uint64) ([]DataSignatures.GetProduct, error) {
	db := productQuery.dbClient.GetDB()

	query, err := db.Prepare(`SELECT p.product_id, p.name, p.color, p.price, p.weight, p.quantity
									FROM Product AS p
									INNER JOIN WatchList AS w ON p.product_id = w.product_id
									WHERE w.account_id = $1 AND w.product_id = $2`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	row, err := query.Query(accountID, productID)

	if err != nil {
		log.Fatal(err)
	}

	var products []DataSignatures.GetProduct
	for row.Next() {
		var product DataSignatures.GetProduct
		err = row.Scan(&product.Id, &product.Name, &product.Color, &product.Price, &product.Weight, &product.Quantity)

		if err != nil {
			log.Fatal(err)
		}

		products = append(products, product)
	}

	return products, nil
}

func (productQuery *ProductQuery) GetProductsOfCartList(accountID uint64, productID uint64) ([]DataSignatures.GetProductFromCart, error) {
	db := productQuery.dbClient.GetDB()

	query, err := db.Prepare(`SELECT p.product_id, p.name, p.color, p.price, p.weight, p.quantity, c.product_count
									FROM Product AS p
									INNER JOIN Cart AS c ON p.product_id = c.product_id
									WHERE c.account_id = $1 AND c.product_id = $2`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	row, err := query.Query(accountID, productID)

	if err != nil {
		log.Fatal(err)
	}

	var products []DataSignatures.GetProductFromCart
	for row.Next() {
		var product DataSignatures.GetProductFromCart
		err = row.Scan(&product.Id, &product.Name, &product.Color, &product.Price,
			&product.Weight, &product.Quantity, &product.RequestedQuantity)

		if err != nil {
			log.Fatal(err)
		}

		products = append(products, product)
	}

	return products, nil
}

// -------------------------------- POST ----------------------------------

func (productQuery *ProductQuery) AddProductToOrder(orderID uint64, productID uint64, productCount uint64) error {
	db := productQuery.dbClient.GetDB()

	query, err := db.Prepare(`INSERT INTO order_product (order_id, product_id, product_count)
									SELECT *
										FROM (
											VALUES ($1, $2, $3)
											 ) AS input
									WHERE (
										SELECT quantity
										FROM product
										WHERE product_id = $2
											  ) >= $3;`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	_, err = query.Exec(orderID, productID, productCount)

	if err != nil {
		return err
	}

	return nil
}
