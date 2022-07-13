package Handlers

import (
	"API/Database"
	"API/Communication/DataSignatures"
	q "API/Infrastructure/Queries"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	dbClient *Database.Postgresql
}

func (productHandler *ProductHandler) NewProductHandler(dbClient *Database.Postgresql) *ProductHandler {
	return &ProductHandler{dbClient: dbClient}
}

func (productHandler *ProductHandler) ProductByCategoryNameHandler(c *gin.Context) {
	var categ DataSignatures.PostCategory
	if err := c.ShouldBindJSON(&categ); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"err": "invalid json",
		})
		return
	}
	productq := q.NewProductQuery(productHandler.dbClient)
	prods, err := productq.GetProductsByCategoryName(categ.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	response := make([]map[string]interface{}, 0, 0)
	for i := 0; i < len(prods); i++ {
		prodObj := make(map[string]interface{})
		prodObj["id"] = prods[i].Id
		prodObj["name"] = prods[i].Name
		prodObj["color"] = prods[i].Color
		prodObj["price"] = prods[i].Price
		prodObj["weight"] = prods[i].Weight
		prodObj["quantity"] = prods[i].Quantity

		response = append(response, prodObj)
	}

	c.JSON(http.StatusOK, response)

}