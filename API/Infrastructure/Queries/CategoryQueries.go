package Queries

import (
	"API/Communication/DataSignatures"
	"API/Database"
	"log"
)

type CategoryQuery struct {
	dbClient *Database.Postgresql
}

func NewCategoryQuery(dbClient *Database.Postgresql) *CategoryQuery {
	return &CategoryQuery{dbClient: dbClient}
}

func (categoryQuery *CategoryQuery) GetAllCategories() ([]DataSignatures.Category, error) {
	db := categoryQuery.dbClient.GetDB()

	query, err := db.Prepare(`SELECT category_id, name
    								FROM category`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	row, err := query.Query()

	if err != nil {
		log.Fatal(err)
	}

	var categories []DataSignatures.Category
	for row.Next() {
		var category DataSignatures.Category

		err = row.Scan(&category.Id, &category.Name)

		if err != nil {
			log.Fatal(err)
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (categoryQuery *CategoryQuery) AddNewCategory(category *DataSignatures.Category) error {
	db := categoryQuery.dbClient.GetDB()

	query, err := db.Prepare(`INSERT INTO Category (name)
									VALUES ($1)`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	_, err = query.Exec(category.Name)

	if err != nil {
		return err
	}

	return nil
}
