package Queries

import (
	"API/Communication/DataSignatures"
	"API/Database"
	"log"
)

type OrderQuery struct {
	dbClient *Database.Postgresql
}

func NewOrderQuery(dbClient *Database.Postgresql) *OrderQuery {
	return &OrderQuery{dbClient: dbClient}
}

// -------------------------------- GET ----------------------------------

func (orderQuery *OrderQuery) GetOrdersByAccountID(id uint64) ([]DataSignatures.GetOrder, error) {
	db := orderQuery.dbClient.GetDB()

	query, err := db.Prepare(`SELECT order_id, description, delivery_method, order_date, address
									FROM orderitem 
									WHERE account_id = $1`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	row, err := query.Query(id)

	if err != nil {
		log.Fatal(err)
	}

	var orders []DataSignatures.GetOrder
	for row.Next() {
		var order DataSignatures.GetOrder

		err = row.Scan(&order.Id, &order.Description, &order.DeliveryMethod, &order.OrderDate, &order.Address)

		if err != nil {
			log.Fatal(err)
		}

		orders = append(orders, order)
	}

	return orders, nil
}

// -------------------------------- POST ----------------------------------

func (orderQuery *OrderQuery) CreateOrder(accountID uint64, order DataSignatures.PostOrder) error {
	db := orderQuery.dbClient.GetDB()

	query, err := db.Prepare(`CALL CreateOrderAndClearCart($1, $2, $3, $4, $5)`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	_, err = query.Exec(accountID, order.Description, order.Address, order.DeliveryMethod, order.OrderDate)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
