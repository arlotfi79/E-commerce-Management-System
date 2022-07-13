package Handlers

import (
	"API/Database"
	q "API/Infrastructure/Queries"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	dbClient *Database.Postgresql
}

func (categoryHandler *CategoryHandler) NewCategoryHandler(dbClient *Database.Postgresql) *CategoryHandler {
	return &CategoryHandler{dbClient: dbClient}
}

func (categoryHandler *CategoryHandler) GetAllCategoriesHandler(c *gin.Context) {
	
	catq := q.NewCategoryQuery(categoryHandler.dbClient)
	categs, err := catq.GetAllCategories()
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	response := make([]map[string]interface{}, 0, 0)
	for i := 0; i < len(categs); i++ {
		categObj := make(map[string]interface{})
		categObj["id"] = categs[i].Id
		categObj["name"] = categs[i].Name
		response = append(response, categObj)
	}

	c.JSON(http.StatusOK, response)
}
