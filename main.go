package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/rralcala/martyn/lib/log"
	"github.com/rralcala/martyn/models"
)

func parseJSONArray(arrayQuery string) []string {
	var sort []string

	if len(arrayQuery) > 0 {
		unmarshaled := json.Unmarshal([]byte(arrayQuery), &sort)
		if unmarshaled != nil {
			log.Warning(fmt.Sprintf("Array parameter error: %s", arrayQuery))
		} else {
			return sort
		}
	}
	return nil

}

func parseJSONArrayInt(arrayQuery string) []int {
	var sort []int

	if len(arrayQuery) > 0 {
		unmarshaled := json.Unmarshal([]byte(arrayQuery), &sort)
		if unmarshaled == nil {
			return sort
		}
		log.Warning(fmt.Sprintf("Array parameter error: %s", arrayQuery))
	}
	return nil

}

func parseJSONMap(sortQuery string) map[string]interface{} {
	var sort map[string]interface{}

	if len(sortQuery) > 0 {
		unmarshaled := json.Unmarshal([]byte(sortQuery), &sort)
		if unmarshaled != nil {
			log.Error("Map parameter error")
		} else {
			return sort
		}
	}
	return nil

}

// CORSMiddleware Allows all origins for now
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "X-Total-Count")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	models.ConnectDatabase()
	router.GET("/transactions", getTransactions)
	router.GET("/transactions/:id", getTransactionByID)
	router.POST("/transactions", postTransaction)
	router.DELETE("/transactions", deleteTransactions)
	router.DELETE("/transactions/:id", deleteTransactions)
	router.PUT("/transactions/:id", updateTransaction)
	router.PUT("/transactions", updateTransaction)

	router.GET("/accounts", getAccounts)
	router.GET("/accounts/:id", getAccountByID)
	router.POST("/accounts", postAccount)
	router.DELETE("/accounts", deleteAccounts)
	router.DELETE("/accounts/:id", deleteAccounts)
	router.PUT("/accounts/:id", updateAccounts)
	router.PUT("/accounts", updateAccounts)

	router.GET("/providers", getProviders)
	router.GET("/providers/:id", getProviderByID)
	router.POST("/providers", postProvider)
	router.DELETE("/providers", deleteProviders)
	router.DELETE("/providers/:id", deleteProviders)
	router.PUT("/providers/:id", updateProviders)
	router.PUT("/providers", updateProviders)

	router.GET("/cost-centers", getCostCenters)
	router.GET("/cost-centers/:id", getCostCenterByID)
	router.POST("/cost-centers", postCostCenter)
	router.DELETE("/cost-centers", deleteCostCenters)
	router.DELETE("/cost-centers/:id", deleteCostCenters)
	router.PUT("/cost-centers/:id", updateCostCenters)
	router.PUT("/cost-centers", updateCostCenters)

	router.Run("localhost:8080")
}
