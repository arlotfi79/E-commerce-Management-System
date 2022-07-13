package Queries

import (
	"API/Communication/DataSignatures"
	"API/Database"
	"log"
	"time"
)

type PromotionCodeQuery struct {
	dbClient *Database.Postgresql
}

func NewPromotionCodeQuery(dbClient *Database.Postgresql) *PromotionCodeQuery {
	return &PromotionCodeQuery{dbClient: dbClient}
}

func (promotionCodeQuery *PromotionCodeQuery) GetPromotionCodeByID(id uint64) (DataSignatures.PromotionCode, error) {
	db := promotionCodeQuery.dbClient.GetDB()

	query, err := db.Prepare(`SELECT promotion_code_id, value, expire_date
									FROM PromotionCode
									WHERE promotion_code_id = $1`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	row := query.QueryRow(id)

	var promotionCode DataSignatures.PromotionCode
	err = row.Scan(&promotionCode.Id, &promotionCode.Value, &promotionCode.ExpireDate)

	if err != nil {
		log.Fatal(err)
	}

	return promotionCode, nil

}

func (promotionCodeQuery *PromotionCodeQuery) AddNewPromotionCode(value float64, expireDate time.Time) error {
	db := promotionCodeQuery.dbClient.GetDB()

	query, err := db.Prepare(`INSERT INTO PromotionCode (value, expire_date) 
									VALUES ($1, $2, $3)`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	if err != nil {
		log.Fatal(err)
	}

	_, err = query.Exec(value, expireDate)

	if err != nil {
		return err
	}

	return nil

}
