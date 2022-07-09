package Database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type Postgresql struct {
	db *sql.DB
}

func (dbClient *Postgresql) Init() error {
	DbUsername := os.Getenv("DB_USERNAME")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")
	DbName := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		DbUsername, DbName, DbPassword, DbHost, DbPort)
	var err error
	dbClient.db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = dbClient.db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Print(fmt.Sprintf("Connected to PostgreSQL on port 5432"))
	return nil
}

func (dbClient *Postgresql) Close() error {
	log.Print(fmt.Sprintf("Disconnecting from PostgreSQL ..."))
	return dbClient.db.Close()
}

func (dbClient *Postgresql) GetDB() *sql.DB {
	return dbClient.db
}
