package Handlers

import (
	"API/Communication/DataSignatures"
	"API/Database"
	q "API/Infrastructure/Queries"
	"log"
	"net/http"
	"API/Infrastructure/auth"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CartHandler struct {
	dbClient *Database.Postgresql
	authInterface          auth.AuthInterface
	tokenInterface         auth.TokenInterface
}

func (cartHandler *CartHandler) NewCartHandler(dbClient *Database.Postgresql, authInt auth.AuthInterface, tokenInt auth.TokenInterface) *CartHandler {
	return &CartHandler{dbClient: dbClient, authInterface: authInt, tokenInterface: tokenInt}
}


func (cartHandler *CartHandler) AddToCartHandler (c *gin.Context) {
	var cart DataSignatures.PostCart
	accessInfo, err := cartHandler.tokenInterface.ExtractTokenMetadata(c.Request)
  if err == nil {
		if err := c.ShouldBindJSON(&cart); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"err": "invalid json",
			})
			return
		}
		cartq := q.NewCartQuery(cartHandler.dbClient)
  	err = cartq.AddProductToCartList(accessInfo.UserId, cart.ProductId, cart.ProductCount)
		if err != nil {
			c.JSON(http.StatusNotFound, err)
			return
		}
		return 
	} else {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}
}

func (cartHandler *CartHandler) GetCartHandler (c *gin.Context) {
	// var cart DataSignatures.PostCart
	accessInfo, err := cartHandler.tokenInterface.ExtractTokenMetadata(c.Request)
  if err == nil {
		// if err := c.ShouldBindJSON(&cart); err != nil {
		// 	c.JSON(http.StatusUnprocessableEntity, gin.H{
		// 		"err": "invalid json",
		// 	})
		// 	return
		// }
		promoCode, _ := strconv.ParseUint(c.Query("promoCode"), 10, 32)
		
		log.Println(promoCode)

		var prods []DataSignatures.GetProductFromCart
		prodq := q.NewProductQuery(cartHandler.dbClient)
		if promoCode != 0 {
			promoq := q.NewPromotionCodeQuery(cartHandler.dbClient)
			promo, _ := promoq.GetPromotionCodeByID(uint64(promoCode))
			
			prods,err = prodq.GetProductsOfCartListWithTotalCostAndDeductedPromotionCodeByAccountID(accessInfo.UserId, promo.Value)
		} else {
			prods,err = prodq.GetProductsOfCartListWithTotalCostAndDeductedPromotionCodeByAccountID(accessInfo.UserId, 0)
		}

		c.JSON(http.StatusOK, prods)
		return 
	} else {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}
}
