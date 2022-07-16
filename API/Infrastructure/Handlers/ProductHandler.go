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

type ProductHandler struct {
	dbClient       *Database.Postgresql
	authInterface  auth.AuthInterface
	tokenInterface auth.TokenInterface
}

func (productHandler *ProductHandler) NewProductHandler(dbClient *Database.Postgresql, authInt auth.AuthInterface, tokenInt auth.TokenInterface) *ProductHandler {
	return &ProductHandler{dbClient: dbClient, authInterface: authInt, tokenInterface: tokenInt}
}

func (productHandler *ProductHandler) ProductByCategoryNameHandler(c *gin.Context) {
	var categ DataSignatures.GetCategory
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

func (productHandler *ProductHandler) AddNewProductHandler(c *gin.Context) {
	_, err := productHandler.tokenInterface.ExtractTokenMetadata(c.Request)

	if err == nil {
		var product DataSignatures.PostProduct
		if err := c.ShouldBindJSON(&product); err != nil {
			log.Fatal(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"err": "invalid json",
			})
			return
		}

		reviewq := q.NewProductQuery(productHandler.dbClient)

		err := reviewq.AddNewProduct(&product)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
	}
}
