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

	query, err := db.Prepare(`INSERT INTO Account (name, last_name, user_name, phone_number, password, email, gender, birth_date, join_date)
									VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`)

	if err != nil {
		log.Fatal(err)
		return err
	}

	defer query.Close()

	_, err = query.Exec(
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
		return err
	}

	return nil
}
