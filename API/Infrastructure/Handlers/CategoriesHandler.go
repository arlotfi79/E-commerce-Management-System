package Handlers

import (
	"API/Communication/DataSignatures"
	"API/Database"
	q "API/Infrastructure/Queries"
	"API/Infrastructure/auth"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type CategoryHandler struct {
	dbClient       *Database.Postgresql
	authInterface  auth.AuthInterface
	tokenInterface auth.TokenInterface
}

func (categoryHandler *CategoryHandler) NewCategoryHandler(dbClient *Database.Postgresql, authInt auth.AuthInterface, tokenInt auth.TokenInterface) *CategoryHandler {
	return &CategoryHandler{dbClient: dbClient, authInterface: authInt, tokenInterface: tokenInt}
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

func (categoryHandler *CategoryHandler) AddNewCategoryHandler(c *gin.Context) {
	_, err := categoryHandler.tokenInterface.ExtractTokenMetadata(c.Request)

	if err == nil {
		var category DataSignatures.Category
		if err := c.ShouldBindJSON(&category); err != nil {
			log.Fatal(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"err": "invalid json",
			})
			return
		}

		categoryq := q.NewCategoryQuery(categoryHandler.dbClient)

		err := categoryq.AddNewCategory(&category)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
	}
}
