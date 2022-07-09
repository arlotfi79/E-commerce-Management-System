package Queries

import (
	"DBProject/API/Communication/DataSignatures"
	"DBProject/API/Database"
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

	query, err := db.Prepare("INSERT" +
		"Account" +
		"VALUES" +
		"(?, ?, ?, ?, ?, ?, ?, ?, ?, ?")

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	row, err := query.Query(account.NationalCode,
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
