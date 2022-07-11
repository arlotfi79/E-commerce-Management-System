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

func (UserQuery *UserQuery) CreateUser(account *DataSignatures.PostAccount) error {
	db := UserQuery.dbClient.GetDB()

	query, err := db.Prepare(`INSERT INTO Account (name, last_name, user_name, phone_number, password, email, gender, birth_date, join_date)
									VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`)

	if err != nil {
		log.Fatalln(err)
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
		log.Fatalln(err)
		return err
	}

	return nil
}

func (UserQuery *UserQuery) GetUserByUname(username string) ([]DataSignatures.GetAccount, error) {
	db := UserQuery.dbClient.GetDB()

	query, err := db.Prepare(`SELECT * FROM Account WHERE user_name=$1`)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	defer query.Close()
	rows, err := query.Query(username)
	if err != nil {
		log.Fatalln(err)
		return nil, err

	}
	var accounts []DataSignatures.GetAccount
	for rows.Next() {
		var account DataSignatures.GetAccount

		err = rows.Scan(&account.Id, &account.Name, &account.LastName, &account.UserName, &account.PhoneNumber, &account.Password, &account.Email, &account.Gender, &account.BirthDate, &account.JoinDate)

		if err != nil {
			log.Fatalln(err)
		}

		accounts = append(accounts, account)
	}
	return accounts, err
}

func (UserQuery *UserQuery) GetUserByEmail(email string) ([]DataSignatures.GetAccount, error) {
	db := UserQuery.dbClient.GetDB()

	query, err := db.Prepare(`SELECT * FROM Account WHERE email=$1`)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	defer query.Close()
	rows, err := query.Query(email)
	if err != nil {
		log.Fatalln(err)
		return nil, err

	}
	var accounts []DataSignatures.GetAccount
	for rows.Next() {
		var account DataSignatures.GetAccount

		err = rows.Scan(&account.Id, &account.Name, &account.LastName, &account.UserName, &account.PhoneNumber, &account.Password, &account.Email, &account.Gender, &account.BirthDate, &account.JoinDate)

		if err != nil {
			log.Fatalln(err)
		}

		accounts = append(accounts, account)
	}
	return accounts, err
}

func (UserQuery *UserQuery) GetUserByPhoneNumber(phoneNumber string) ([]DataSignatures.GetAccount, error) {
	db := UserQuery.dbClient.GetDB()

	query, err := db.Prepare(`SELECT * FROM Account WHERE phone_number=$1`)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	defer query.Close()
	rows, err := query.Query(phoneNumber)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	var accounts []DataSignatures.GetAccount
	for rows.Next() {
		var account DataSignatures.GetAccount

		err = rows.Scan(&account.Id, &account.Name, &account.LastName, &account.UserName, &account.PhoneNumber, &account.Password, &account.Email, &account.Gender, &account.BirthDate, &account.JoinDate)

		if err != nil {
			log.Fatalln(err)
		}

		accounts = append(accounts, account)
	}
	return accounts, err
}
