package main

import (
	"API/Database"
	"log"
)

func main() {
	var db Database.Postgresql
	err := db.Init()
	defer db.Close()

	if err != nil {
		log.Fatalln(err.Error())
	}
}
