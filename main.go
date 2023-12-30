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

	sort := parseJSONArray(c.Query("sort"))
	itemRange := parseJSONArrayInt(c.Query("range"))
	filter := parseJSONMap(c.Query("filter"))
	transactions := models.GetTransacions(sort, itemRange, filter)
	var ret []models.TransactionOutput
	for _, a := range transactions {
		ret = append(ret, models.Flatten(&a))
	}
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
		if unmarshaled != nil {
			log.Warning(fmt.Sprintf("Array parameter error: %s", arrayQuery))
		} else {
			return sort
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
	c.JSON(http.StatusOK, gin.H{"message": "success"})
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

func main() {
	router := gin.Default()

	models.ConnectDatabase()
	router.GET("/transactions", getTransactions)
	router.POST("/transactions", postTransaction)
	router.GET("/transactions/:id", getTransactionByID)
	router.PUT("/transactions/:id", updateTransaction)
	router.PATCH("/transactions/:id", updateTransaction)
	router.Run("localhost:8080")
}
