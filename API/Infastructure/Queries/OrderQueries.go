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

func (orderQuery *OrderQuery) GetOrdersByAccountID(id string) ([]DataSignatures.Order, error) {
	db := orderQuery.dbClient.GetDB()

	query, err := db.Prepare(`SELECT order_id, description, delivery_method, order_date, address
									FROM orderitem 
									WHERE national_code = $1`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	row, err := query.Query(id)

	if err != nil {
		log.Fatal(err)
	}

	var orders []DataSignatures.Order
	for row.Next() {
		var order DataSignatures.Order

		err = row.Scan(&order.Id, &order.Description, &order.DeliveryMethod, &order.OrderDate, &order.Address)

		if err != nil {
			log.Fatal(err)
		}

		orders = append(orders, order)
	}

	return orders, nil
}

// -------------------------------- POST ----------------------------------
