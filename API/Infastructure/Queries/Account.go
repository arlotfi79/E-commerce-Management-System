package Queries

import (
	"API/Communication/DataSignatures"
	"API/Database"
	"log"
)

type UserQuery struct {
	dbClient *Database.Postgresql
}

func NewUserQuery(dbClient *Database.Postgresql) *UserQuery {
	return &UserQuery{dbClient: dbClient}
}

func (UserQuery *UserQuery) CreateUser(account *DataSignatures.Account) error {
	db := UserQuery.dbClient.GetDB()

	query, err := db.Prepare("INSERT " +
		"INSERT " +
		"Account " +
		"VALUES " +
		"($1, $2, $3, $4, $5, $6, $7, $8, $9, $10")

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	_, err = query.Exec(account.NationalCode,
		account.Name,
		account.LastName,
		account.UserName,
		account.PhoneNumber,
		account.Password,
		account.Email,
		account.Gender,
		account.BirthDate,
		account.JoinDate,
	)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
