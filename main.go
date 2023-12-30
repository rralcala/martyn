package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/rralcala/martyn/lib/log"
	"github.com/rralcala/martyn/models"
)

func getTransactions(c *gin.Context) {
	count := models.TotalTransactions()
	sort := parseJSONArray(c.Query("sort"))
	itemRange := parseJSONArrayInt(c.Query("range"))
	filter := parseJSONMap(c.Query("filter"))
	transactions := models.GetTransacions(sort, itemRange, filter)
	var ret []models.TransactionOutput
	for _, a := range transactions {
		ret = append(ret, models.Flatten(&a))
	}
	c.Writer.Header().Set("X-Total-Count", strconv.FormatInt(count, 10))
	c.IndentedJSON(http.StatusOK, ret)
}

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
		} else {
			log.Warning(fmt.Sprintf("Array parameter error: %s", arrayQuery))
		}
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

// getTransactionByID locates the transaction whose ID value matches the id
// parameter sent by the client, then returns that transaction as a response.
func getTransactionByID(c *gin.Context) {
	log.Info("getTransactionByID")
	// Loop over the list of transactions, looking for
	// a transaction whose ID value matches the parameter.
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id is not int"})
	}
	transaction, err := models.GetTransactionByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "transaction not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, models.Flatten(transaction))
}

func updateTransaction(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{"message": "success"})
}

func deleteTransactions(c *gin.Context) {
	if len(c.Param("id")) > 0 {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "id is not int"})
			return
		}
		item := models.Transaction{
			ID: id,
		}
		models.DeleteTransactions([]*models.Transaction{&item})
	} else if filter := c.Query("filter"); len(filter) > 0 {
		var items []*models.Transaction
		for _, id := range parseJSONArrayInt(filter) {
			items = append(items, &models.Transaction{
				ID: int64(id),
			})
		}
		models.DeleteTransactions(items)

	}
	c.JSON(http.StatusNoContent, gin.H{"message": "success"})
}

// postTransaction adds a transaction from JSON received in the request body.
func postTransaction(c *gin.Context) {
	var newTransaction models.TransactionInput

	// Call BindJSON to bind the received JSON to
	// newTransaction.
	if err := c.BindJSON(&newTransaction); err != nil {
		return
	}

	// Add the new album to the slice.
	transactionStruct, err := models.Build(&newTransaction)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	models.AppendTransacions(transactionStruct)

	c.IndentedJSON(http.StatusCreated, newTransaction)
}

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
	router.Run("localhost:8080")
}
